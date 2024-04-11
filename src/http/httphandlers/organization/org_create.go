package organization

import (
	"net/http"
	"time"

	"github.com/golang/glog"
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
	"lizisky.com/lizisky/src/utils/orgutil"
)

//
// Create new Organization
//

// define http request param
type requestParam_createOrganization struct {
	// SessionToken string               `json:"sessionToken,omitempty"`                 // session token
	OrgInfo orgtype.Organization `json:"orginfo,omitempty" validate:"omitempty"` // org info
}

// validate request parameter
func (param *requestParam_createOrganization) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

// define http request response
type requestResponse_createOrganization struct {
	OrgInfo *orgtype.Organization `json:"orginfo,omitempty" validate:"omitempty"` // org info
	Staff   *orgtype.OrgStaff     `json:"staff,omitempty"`                        // 需要加入一个组织的staff info
}

// requestHandler_createOrganization implements the HttpHandler interface
type requestHandler_createOrganization struct {
}

func (handler *requestHandler_createOrganization) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_createOrganization) URL() string {
	return httpurl.Url_Org_Create_Organization
}

func (handler *requestHandler_createOrganization) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_createOrganization
	return &param
}

// Handle 处理请求
func (handler *requestHandler_createOrganization) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_createOrganization)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return nil, httpdata.RcSessionExpired
	// }

	return handler.createOrg(rdata.Magic.SessionInfo, &param.OrgInfo)
}

// --------------------------------------------------------------------------------------------------------
// createOrg 创建一个新的组织
func (handler *requestHandler_createOrganization) createOrg(sinfo *sessiondata.SessionInfo, orginfo *orgtype.Organization) (interface{}, httpdata.HttpError) {
	// firstly, check category ID, MUST set major category ID
	if !orgutil.CheckOrgCategory(orginfo.MajorCategory, orginfo.SubCategory) {
		return nil, httpdata.RcInvalidOrgCategory.WithArgs(orginfo.MajorCategory, orginfo.SubCategory)
	}

	// check existing of parent org
	if !orgutil.CheckOrgHierarchic(orginfo.ParentOid, orginfo.RootOid) {
		return nil, httpdata.RcInvalidOrgParentAndRoot.WithArgs(orginfo.ParentOid, orginfo.RootOid)
	}

	orginfo.Oid = 0 // 0 表示创建一个新的组织，忽略客户端传入的 oid
	orginfo.CreateDate = time.Now().UnixMilli()
	orginfo.Creator = sinfo.Aid
	orginfo.Owner = sinfo.Aid
	orginfo.State = orgtype.StateNormal

	if (orginfo.Images != nil) && !orginfo.Images.IsValid() {
		orginfo.Images = nil
	}

	err := dbpool.AddNewRecord(orginfo)
	if err != nil {
		// 写入 DB 失败
		glog.Info("write into DB error", err)
		return nil, httpdata.RcSyetemErr.WithArgs(httpdata.RC_Write2DBFailed)
	}

	orgutil.BuildDownloadURL_org(orginfo)

	// success
	response := httpdata.NewResponseSucc(requestResponse_createOrganization{
		OrgInfo: orginfo,
		Staff:   handler.addMe_as_staff(orginfo, sinfo),
	})
	return response, nil
}

// --------------------------------------------------------------------------------------------------------
// createOrg 创建一个新的组织
func (handler *requestHandler_createOrganization) addMe_as_staff(org *orgtype.Organization, sinfo *sessiondata.SessionInfo) *orgtype.OrgStaff {

	account, _ := dbpool.GetAccountByAID(sinfo.Aid)
	if account == nil {
		return nil
	}

	staff := &orgtype.OrgStaff{
		Aid:         sinfo.Aid,
		Aliasname:   sinfo.Nickname,
		Mobile:      account.Mobile,
		Oid:         org.Oid,
		ParentOid:   org.ParentOid,
		RootOid:     org.RootOid,
		Role:        orgtype.RoleOwner,
		RoleName:    "创始人",
		AddBy:       sinfo.Aid,
		AddByName:   sinfo.Nickname,
		CreateDate:  time.Now().UnixMilli(),
		Description: "",
	}

	err := dbpool.AddNewRecord(staff)
	if err != nil {
		// 写入 DB 失败
		glog.Info("write into DB error", err)
		return nil
	}

	return staff
}
