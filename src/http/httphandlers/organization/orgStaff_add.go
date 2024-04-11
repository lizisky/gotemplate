package organization

import (
	"net/http"
	"time"

	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
	"lizisky.com/lizisky/src/utils/orgutil"
)

// 机构工作人员操作，添加

// define http request param
type requestParam_addOrgStaff struct {
	Staff orgtype.OrgStaff `json:"staff,omitempty"` // 需要加入一个组织的staff info
}

// validate request parameters
func (param *requestParam_addOrgStaff) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	if (param.Staff.Aid == 0) || (param.Staff.Oid == 0) {
		// aid, oid 都必须有效
		return httpdata.RcInvalidHttpRequestArgs
	}

	return nil
}

// define http request response“
type requestResponse_addOrgStaff struct {
	Staff *orgtype.OrgStaff `json:"staff,omitempty"` // 需要加入一个组织的staff info
}

// requestHandler_addOrgStaff implements the HttpHandler interface
type requestHandler_addOrgStaff struct {
}

func (handler *requestHandler_addOrgStaff) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_addOrgStaff) URL() string {
	return httpurl.Url_Org_add_staff_in_org
}

func (handler *requestHandler_addOrgStaff) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_addOrgStaff
	return &param
}

// Handle http request
func (handler *requestHandler_addOrgStaff) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_addOrgStaff)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	return handler.createStaff(rdata.Magic.SessionInfo, &param.Staff)
}

// ------------------------------------------------------------------------------------------------------------------------------
// createStaff 添加一个员工
func (handler *requestHandler_addOrgStaff) createStaff(sinfo *sessiondata.SessionInfo, staff *orgtype.OrgStaff) (interface{}, httpdata.HttpError) {
	if sinfo.Aid == staff.Aid {
		return nil, httpdata.RcYouCannotAddYourself
	}

	if !orgutil.CheckOrgHierarchic_full(staff.Oid, staff.ParentOid, staff.RootOid) {
		return nil, httpdata.RcInvalidOrgHierarchic.WithArgs(staff.Oid, staff.ParentOid, staff.RootOid)
	}

	if dbpool.IsStaffInOrg(staff.Aid, staff.Oid) {
		// 这个 account 已经是这个 org 的员工了
		return nil, httpdata.RcUserAlreadyStaffOfOrg.WithArgs(staff.Aid, staff.Oid)
	}

	// 这个AID必须存在
	if !dbpool.IsAccountExisting(staff.Aid) {
		return nil, httpdata.RcAccountNotExisting.WithArgs(staff.Aid)
	}

	// 只有 Owner/ViceOwner 有权限添加员工
	err := orgutil.IsUserHasRight(sinfo.Aid, staff.Oid, orgtype.RoleOwner, orgtype.RoleViceOwner)
	if err != nil {
		return nil, err
	}

	staff.ID = 0 // useless
	staff.CreateDate = time.Now().UnixMilli()
	staff.AddBy = sinfo.Aid
	staff.AddByName = sinfo.Nickname

	dbpool.AddNewRecord(staff)

	response := &requestResponse_addOrgStaff{
		Staff: staff,
	}

	return httpdata.NewResponseSucc(response), nil
}
