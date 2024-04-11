package organization

import (
	"net/http"

	"lizisky.com/lizisky/src/basictypes/accounts"
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/accountUtil"
)

// 机构工作人员操作，读取
// aid, oid 都必须有效

// define http request param
type requestParam_readOrgStaff struct {
	// SessionToken string `json:"sessionToken,omitempty"`        // session token
	Aid uint64 `json:"aid,omitempty" validate:"gt=0"` // org info
	Oid uint64 `json:"oid,omitempty" validate:"gt=0"` // org info
}

// validate request parameters
func (param *requestParam_readOrgStaff) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

// define http request response“
type requestResponse_readOrgStaff struct {
	Staff   *orgtype.OrgStaff `json:"staff,omitempty"`   // 一个组织中的 staff info
	Account *accounts.Account `json:"account,omitempty"` // 这个 staff 的详细个人 account info
}

// requestHandler_readOrgStaff implements the HttpHandler interface
type requestHandler_readOrgStaff struct {
}

func (handler *requestHandler_readOrgStaff) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_readOrgStaff) URL() string {
	return httpurl.Url_Org_read_staff_in_org
}

func (handler *requestHandler_readOrgStaff) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_readOrgStaff
	return &param
}

// Handle http request
func (handler *requestHandler_readOrgStaff) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_readOrgStaff)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return nil, httpdata.RcSessionExpired
	// }

	return handler.readStaff(param.Aid, param.Oid)
}

// ------------------------------------------------------------------------------------------------------------------------------
// readStaff 读取一个员工
func (handler *requestHandler_readOrgStaff) readStaff(aid, oid uint64) (interface{}, httpdata.HttpError) {
	read_staff, err := dbpool.GetStaffInOrg(aid, oid)
	if err != nil {
		return nil, httpdata.RcUserIsNotStaffOfOrg.WithArgs(aid, oid)
	}

	accountRead := accountUtil.GetAccountByAID(aid)
	if accountRead == nil {
		return nil, httpdata.RcAccountNotExisting.WithArgs(aid)
	}

	response := httpdata.NewResponseSucc(&requestResponse_readOrgStaff{
		Staff:   read_staff,
		Account: accountRead,
	})
	return response, nil
}
