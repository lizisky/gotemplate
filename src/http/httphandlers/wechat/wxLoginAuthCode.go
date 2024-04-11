package wechat

import (
	"net/http"
	"time"

	"github.com/lizisky/liziutils/utils"
	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/basictypes/constants"
	"lizisky.com/lizisky/src/config"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/sessionmgr"
	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
	"lizisky.com/lizisky/src/utils/accountUtil"
	"lizisky.com/lizisky/src/utils/qiniuUtil"
	"lizisky.com/lizisky/src/utils/wxutil"
)

//
// 用户使用微信登录，获取到 WX AuthCode，使用这个 AuthCode 调用后台，
// 后台从 WX Server 获取 user OpenID & UnionID
//
// 在LiziCS中，要求用户必须输入手机号码，
// 所以在返回给client的AccountInfo中，如果用户没有手机号码，则需要显示输入用户信息的UI，提示用户输入手机号码
//
// ClientVersion: 只有在 mode==1 时，才会检查这个参数，如果这个参数小于 config.MinVersion，则返回错误

// Mode
// 1: default value, return all data
// 2: return only session token
// 3: return only qiniu upload token
const (
	mode_AllData = 1
	mode_Session = 2
	mode_Qiniu   = 3
)

// define http request param
type requestParam_wxAuthCode struct {
	// WeChat Auth code
	// 长度验证，参见：constants.MinLength_wxLoginAuthCode
	AuthCode      string `json:"authCode,omitempty" validate:"omitempty,min=10"`
	Mode          int    `json:"mode,omitempty"`
	ClientVersion uint32 `json:"version,omitempty"`
}

func (acp *requestParam_wxAuthCode) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(acp); err != nil {
	// 	return httpdata.RcSyetemErr.WithArgs(httpdata.RC_InvalidHttpRequestParams)
	// }
	return nil
}

type qiniu_data struct {
	// qiniu upload token, 用于上传用户头像
	// return these data here, it's reduces one http request for client
	UploadToken string `json:"uploadToken,omitempty"` // qiniu upload token
	Expires     uint64 `json:"expires,omitempty"`     // qiniu upload token expires time in milliseconds
}

// 打包数据，这些返回数据，不区分用户类型，都返回
type return_data_general struct {
	SessionToken string            `json:"sessionToken,omitempty"` // session token
	Account      *accounts.Account `json:"account,omitempty"`      // account info
	Qiniu        *qiniu_data       `json:"qiniu,omitempty"`        // qiniu upload token
}

// // 打包数据，这些是返回给 org staff 的数据
// type return_data_for_staff struct {
// 	// organization list for Org Staff, 用户属于这些org里面的工作人员
// 	// OrgListStaff []*orgtype.Organization `json:"orgListStaff,omitempty"`
// 	OrgListStaff []*orgmix.OrgDeepInfo `json:"orgListStaff,omitempty"`
// }

// 打包数据，这些是返回给 org membership 的数据
// type return_data_for_membership struct {
// 	// 下面的列表是用户作为学员的相关信息
// 	IncomingClasses []*eduorg.IncomingClass `json:"incomingClasses,omitempty"` // 后续要参加的上课列表，按照时间
// 	CourseList      []*eduorg.CourseInfo    `json:"courseList,omitempty"`      // AID 报名的所有课程列表

// 	// 1. 对于新注册用户，必须返回一些 org 给他，让他选择
// 	// 2. 对于已经注册的用户，如果用户没有选择过 org，则也需要返回一些 org 给他，让他选择
// 	// 3. 对于已经注册的用户，如果用户是某些机构的工作人员，但是他也可以报名其他机构，作为其他机构的学员
// 	// 4. 对于已经注册的用户，如果用户是某些机构的学员，但是他也可以报名其他机构，作为其他机构的学员
// 	// 这里面的 org，必须不能与 OrgListStaff 重复
// 	OrgListVisitor []*orgtype.Organization `json:"orgListVisitor,omitempty"` // organization list for Org Visitors
// }

// define http request response
type requestResponse_wxAuthCode struct {
	// general data for all users
	GeneralData *return_data_general `json:"generalData,omitempty"`

	// data for org staff
	// StaffData *orgmix.Homepage_data_for_staff `json:"staffData,omitempty"`

	// // data for org membership
	// MembershipData *orgmix.Homepage_data_for_membership `json:"membershipData,omitempty"`
}

// requestHandler_wxLoginAuthCode implements the HttpHandler interface
type requestHandler_wxLoginAuthCode struct {
}

func (handler *requestHandler_wxLoginAuthCode) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_wxLoginAuthCode) URL() string {
	return httpurl.Url_Wechat_Login_Auth_Code
}

func (handler *requestHandler_wxLoginAuthCode) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_wxAuthCode
	return &param
}

func (handler *requestHandler_wxLoginAuthCode) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	params, ok := rdata.Data.(*requestParam_wxAuthCode)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	cfg := config.GetConfig()
	if params.ClientVersion < cfg.MinVersion {
		return nil, httpdata.RcClientVersionExpired
	}

	return handler.processLoginAuthCode(params)

}

// processLoginAuthCode
func (handler *requestHandler_wxLoginAuthCode) processLoginAuthCode(params *requestParam_wxAuthCode) (interface{}, httpdata.HttpError) {
	// fmt.Println("===1111====", utils.ToJSONIndent(params))
	wx_session_key, wx_openid, err := wxutil.GetWXopenID(params.AuthCode)
	// fmt.Println("---kvn---parse wechat session info", wx_session_key, wx_openid, err)
	if err != nil {
		return nil, httpdata.NewError(httpdata.RC_GetWechatOpenidError, err)
	}
	if len(wx_openid) < constants.MinLength_wxOpenID {
		return nil, httpdata.RcGetWechatOpenidError
	}

	switch params.Mode {
	case mode_AllData:
		// return nil, httpdata.RcSessionExpired
		return handler.process_all_data(wx_session_key, wx_openid)
	case mode_Session:
		return handler.get_session_token_only(wx_session_key, wx_openid)
	case mode_Qiniu:
		return handler.get_qiniu_token_only(wx_session_key, wx_openid)
	default:
		return nil, httpdata.RcInvalidHttpRequestArgs
	}
}

// get session token only
func (handler *requestHandler_wxLoginAuthCode) get_session_token_only(wx_session_key, wx_openid string) (interface{}, httpdata.HttpError) {
	act := accountUtil.GetAccountByWXopenid(wx_openid)
	if act == nil {
		return nil, httpdata.RcAccountNotExisting.WithArgs("")
	}

	// add to Session Manager
	sessionInfo := &sessiondata.SessionInfo{
		SessionKey:     utils.NewRandomUUID(),
		Aid:            act.Aid,
		Nickname:       act.Nickname,
		AddTime:        uint64(time.Now().UnixMilli()),
		LastAccessTime: uint64(time.Now().UnixMilli()),
		// WxSessionKey:   wx_session_key,
	}
	sessionmgr.AddSession(sessionInfo.SessionKey, sessionInfo)

	responseData := &requestResponse_wxAuthCode{
		GeneralData: &return_data_general{
			SessionToken: sessionInfo.SessionKey,
		},
	}

	return httpdata.NewResponseSucc(responseData), nil
}

// get qiniu upload token only
func (handler *requestHandler_wxLoginAuthCode) get_qiniu_token_only(wx_session_key, wx_openid string) (interface{}, httpdata.HttpError) {
	uploadToken, expires := qiniuUtil.BuildUploadToken()
	qiniuD := &qiniu_data{
		UploadToken: uploadToken,
		Expires:     expires,
	}

	responseData := &requestResponse_wxAuthCode{
		GeneralData: &return_data_general{
			Qiniu: qiniuD,
		},
	}

	return httpdata.NewResponseSucc(responseData), nil
}

// process all data
func (handler *requestHandler_wxLoginAuthCode) process_all_data(wx_session_key, wx_openid string) (interface{}, httpdata.HttpError) {

	var responseData *requestResponse_wxAuthCode = &requestResponse_wxAuthCode{}

	// first, check is this WeChat account existing in database or not
	act := accountUtil.GetAccountByWXopenid(wx_openid)

	// 1, if act == nil, this is a new account for lizics
	// 2, if he/she doesn't have mobile number,
	//    he/she did not finish the registration process, so we treat he/she as new account
	// is_new_account := (act == nil) || !act.HasMboile()

	if act == nil {
		// right now, this is a new WeChat user
		act = &accounts.Account{
			WxOpenID:   wx_openid,
			CreateDate: time.Now().UnixMilli(),
		}

		err := dbpool.AddNewRecord(act)
		if err != nil {
			// 写入 DB 失败
			return nil, httpdata.RcSyetemErr.WithArgs(httpdata.RC_Write2DBFailed)
		}

		// for new user, we need to return some orgs for him to choose
		// responseData.MembershipData = build_data_for_membership(act.Aid, true)
		// } else if !act.HasMboile() {
		// for existing users, but without mobile number,
		// he/she did not finish the registration process, so we treat he/she as new user
		// we need to return some orgs for him to choose
		// responseData.MembershipData = build_data_for_membership(act.Aid, true)
	} else {
		// for existing users,

		// first, check is he/she a staff of any org
		// responseData.StaffData = orgDeepUtil.ReadHomepageData_For_Staff(act.Aid)
		// second, check is he/she a member of any org
		// responseData.MembershipData = build_data_for_membership(act.Aid, false)
	}

	// responseData.MembershipData = orgDeepUtil.ReadHomepageData_For_Membership(act.Aid, is_new_account)
	responseData.GeneralData = build_general_data(act)

	{
		// fill org list for visitor
		// 目前阶段，仅支持用户一种身份，要么是工作人员，要么是学员
		// 同时支持两种身份的情况，以后再说

		// var excludeOids []uint64
		// if (responseData.StaffData != nil) && (responseData.StaffData.OrgListStaff != nil) && (len(responseData.StaffData.OrgListStaff) > 0) {
		// 	excludeOids = make([]uint64, 0, len(responseData.StaffData.OrgListStaff))
		// 	for idx := len(responseData.StaffData.OrgListStaff) - 1; idx >= 0; idx-- {
		// 		excludeOids = append(excludeOids, responseData.StaffData.OrgListStaff[idx].Oid)
		// 	}
		// }

		// responseData.MembershipData.OrgListVisitor = orgutil.GetOrgList(uint64(0), uint32(0), uint32(0), uint32(0), excludeOids)
	}

	// add to Session Manager
	sessionInfo := &sessiondata.SessionInfo{
		SessionKey:     responseData.GeneralData.SessionToken,
		Aid:            act.Aid,
		Nickname:       act.Nickname,
		AddTime:        uint64(time.Now().UnixMilli()),
		LastAccessTime: uint64(time.Now().UnixMilli()),
		// WxSessionKey:   wx_session_key,
	}
	sessionmgr.AddSession(sessionInfo.SessionKey, sessionInfo)

	return httpdata.NewResponseSucc(responseData), nil
}

// build general data
func build_general_data(account *accounts.Account) *return_data_general {
	uploadToken, expires := qiniuUtil.BuildUploadToken()
	qiniuD := &qiniu_data{
		UploadToken: uploadToken,
		Expires:     expires,
	}

	data := &return_data_general{
		SessionToken: utils.NewRandomUUID(),
		Account:      account,
		Qiniu:        qiniuD,
	}

	return data
}

// // build general data
// func build_data_for_staff(aid uint64) *orgmix.Homepage_data_for_staff {
// 	// 如果当前用户属于机构工作人员，则返回用户所属机构 list 信息
// 	// 用户属于这些org里面的工作人员
// 	orgListStaff := eduorgutil.GetOrgListByStaffID(aid)

// 	if len(orgListStaff) < 1 {
// 		return nil
// 	}

// 	{
// 		// 读取所属的第一个机构的详细信息 ONLY !!!
// 		oid := orgListStaff[0].Oid
// 		orgDeepInfo, _ := orgDeepUtil.ReadOrgDeepInfo(oid, 1)
// 		// orglist_deep := make([]interface{}, 0, len(orgListStaff))
// 		// 第一个包含一个元素的 []interface{}, 里面的元素是 orgDeepInfo
// 		// orglist_deep := make([]interface{}, 1)
// 		orglist_deep := make([]*orgmix.OrgDeepInfo, 1, 1)
// 		orglist_deep[0] = orgDeepInfo
// 		data := &orgmix.Homepage_data_for_staff{
// 			// OrgListStaff: orgListStaff,
// 			OrgListStaff: orglist_deep, //[:len(orglist_deep)],
// 		}

// 		return data
// 	}

// 	// {
// 	// 	// 读取全部所属机构的详细信息
// 	// 	orglist_deep := make([]interface{}, 0, len(orgListStaff))
// 	// 	for _, org := range orgListStaff {
// 	// 		orgDeepInfo, _ := orgDeepUtil.ReadOrgDeepInfo(org.Oid, 1)
// 	// 		if orgDeepInfo != nil {
// 	// 			orglist_deep = append(orglist_deep, orgDeepInfo)
// 	// 		}
// 	// 	}
// 	// 	data := &return_data_for_staff{
// 	// 		// OrgListStaff: orgListStaff,
// 	// 		OrgListStaff: orglist_deep, //[:len(orglist_deep)],
// 	// 	}

// 	// 	return data
// 	// }

// }

// // build general org membership
// func build_data_for_membership(aid uint64, new_registered_user bool) *return_data_for_membership {
// 	// 如果当前用户属于学员，则返回用户所属机构 list 信息
// 	// 用户后续要参加的上课列表，从当前时间开始，获取后续的3次的上课列表
// 	var incomingClasses []*eduorg.IncomingClass = nil
// 	// AID 报名的所有课程列表
// 	// 报名的所有课程列表
// 	var courseList []*eduorg.CourseInfo = nil

// 	// 如果是新注册用户，则不需要返回这些信息, 因为新注册用户还没有报名任何课程
// 	if !new_registered_user {
// 		// 读取用户报名的课程列表, 最多允许欠费的课程数量为10个
// 		courseList = eduorgutil.GetCourseListByMemberID(aid, -10)
// 		// fmt.Println("courseList = ", utils.ToJSONIndent(courseList))

// 		// firstly, read all classes for this account in today
// 		startTime := csTimeUtil.BeginOfToday()
// 		endTime := startTime + uint64(timeUtil.MillisecondPerDay)
// 		incomingClasses = eduorgutil.ReadComingClasses_forAID(aid, 0, startTime, endTime, 10)

// 		// secondly, if no class in today, then read next classes for this account in future
// 		if len(incomingClasses) < 1 {
// 			endTime = 0
// 			incomingClasses = eduorgutil.ReadComingClasses_forAID(aid, 0, startTime, endTime, 1)
// 		}
// 	}

// 	orgListVisitor := orgutil.GetOrgList(uint64(0), uint32(0), uint32(0), uint32(0), nil)

// 	if (len(incomingClasses) < 1) && (len(courseList) < 1) && (len(orgListVisitor) < 1) {
// 		// 如果没有任何数据，则不需要返回这些信息
// 		return nil
// 	}

// 	data := &return_data_for_membership{
// 		CourseList:      courseList,
// 		IncomingClasses: incomingClasses,
// 		OrgListVisitor:  orgListVisitor,
// 	}

// 	return data
// }
