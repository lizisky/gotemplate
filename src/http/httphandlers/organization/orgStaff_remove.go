package organization

import (
	"net/http"

	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/orgutil"
)

// 机构工作人员操作，删除staff

// define http request param
type requestParam_removeOrgStaff struct {
	// SessionToken string `json:"sessionToken,omitempty"`        // session token
	Aid uint64 `json:"aid,omitempty" validate:"gt=0"` // org info
	Oid uint64 `json:"oid,omitempty" validate:"gt=0"` // org info
}

// validate request parameters
func (param *requestParam_removeOrgStaff) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

// define http request response“
// type requestResponse_removeOrgStaff struct {
// 	Staff *orgtype.OrgStaff `json:"staff,omitempty"` // 需要加入一个组织的staff info
// }

// requestHandler_removeOrgStaff implements the HttpHandler interface
type requestHandler_removeOrgStaff struct {
}

func (handler *requestHandler_removeOrgStaff) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_removeOrgStaff) URL() string {
	return httpurl.Url_Org_remove_staff_in_org
}

func (handler *requestHandler_removeOrgStaff) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_removeOrgStaff
	return &param
}

// Handle http request
func (handler *requestHandler_removeOrgStaff) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_removeOrgStaff)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return nil, httpdata.RcSessionExpired
	// }

	if rdata.Magic.SessionInfo.Aid == param.Aid {
		return nil, httpdata.RcYouCannotRemoveYourself
	}

	return handler.deleteStaff(param.Aid, param.Oid)
}

// ------------------------------------------------------------------------------------------------------------------------------
// deleteStaff 删除一个员工
func (handler *requestHandler_removeOrgStaff) deleteStaff(aid, oid uint64) (interface{}, httpdata.HttpError) {
	// 就算org中没有这个员工，也无所谓，我们就当做删除成功了，
	// 所以不需要检查这个员工是否存在
	// if !dbpool.IsStaffInOrg(staff.Aid, staff.Oid, staff.ParentOid) {
	// 	return nil, httpdata.RcUserIsNotStaffOfOrg.WithArgs(staff.Oid, staff.ParentOid, staff.Aid)
	// }

	// 只有 Owner ViceOwner 有权限delete员工
	err := orgutil.IsUserHasRight(aid, oid, orgtype.RoleOwner, orgtype.RoleViceOwner)
	if err != nil {
		return nil, err
	}

	dbpool.RemoveStaffFromOrg(aid, oid)
	return httpdata.NewResponseSucc(), nil
}
