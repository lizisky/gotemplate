package organization

import (
	"net/http"

	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/orgutil"
)

// 机构工作人员操作，更新员工信息
// 1, 员工自己可以修改自己在企业中的昵称和说明信息
// 2，org owner/ViceOwner 可以修改员工的所有信息

// define http request param
type requestParam_updateOrgStaff struct {
	// SessionToken string           `json:"sessionToken,omitempty"` // session token
	Staff orgtype.OrgStaff `json:"staff,omitempty"` // 需要更新的 staff info
}

// validate request parameters
func (param *requestParam_updateOrgStaff) IsValid() httpdata.HttpError {
	if (param.Staff.Aid == 0) || (param.Staff.Oid == 0) {
		// aid, oid 都必须有效
		return httpdata.RcInvalidHttpRequestArgs
	}
	return nil
}

// define http request response“
type requestResponse_updateOrgStaff struct {
	Staff *orgtype.OrgStaff `json:"staff,omitempty"` // 需要加入一个组织的staff info
}

// requestHandler_updateOrgStaff implements the HttpHandler interface
type requestHandler_updateOrgStaff struct {
}

func (handler *requestHandler_updateOrgStaff) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_updateOrgStaff) URL() string {
	return httpurl.Url_Org_update_staff_in_org
}

func (handler *requestHandler_updateOrgStaff) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_updateOrgStaff
	return &param
}

// Handle http request
func (handler *requestHandler_updateOrgStaff) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_updateOrgStaff)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return nil, httpdata.RcSessionExpired
	// }

	return handler.updateStaff(rdata.Magic.SessionInfo.Aid, &param.Staff)
}

// ------------------------------------------------------------------------------------------------------------------------------
// updateStaff 更新一个员工
func (handler *requestHandler_updateOrgStaff) updateStaff(aid uint64, staff *orgtype.OrgStaff) (interface{}, httpdata.HttpError) {
	staffExisting, _ := dbpool.GetStaffInOrg(staff.Aid, staff.Oid)
	if staffExisting == nil {
		// user is not staff of this org
		return nil, httpdata.RcUserIsNotStaffOfOrg.WithArgs(staff.Oid, staff.Aid)
	}

	// change root org is not allowed
	if staff.RootOid != staffExisting.RootOid {
		return nil, httpdata.RcRejectUpdateRootOid.WithArgs(staff.RootOid)
	}
	if staff.ParentOid != staffExisting.ParentOid {
		// if orgutil.CheckOrgHierarchic(staff.ParentOid, staff.RootOid) {
		// 	return nil, httpdata.RcInvalidOrgParentAndRoot.WithArgs(staff.ParentOid, staff.RootOid)
		// }
		if !orgutil.CheckOrgHierarchic_full(staff.Oid, staff.ParentOid, staff.RootOid) {
			return nil, httpdata.RcInvalidOrgHierarchic.WithArgs(staff.Oid, staff.ParentOid, staff.RootOid)
		}
	}

	// 只有 Owner ViceOwner 有权限修改员工信息， 员工自己也可以修改自己的信息
	err := orgutil.IsUserHasRight(aid, staff.Oid, orgtype.RoleOwner, orgtype.RoleViceOwner)
	if err != nil {
		if staffExisting.Aid == aid {
			// 员工自己也可以修改自己的信息
			return handler.updateSelfStaffInfo(staff, staffExisting)
		}

		return nil, err
	}

	staffNew := staffExisting.AssignFields(staff)
	if staffNew == nil {
		// 信息完全相同，没必要 update, do nothing
		return nil, httpdata.RcNeednotUpdateIdenticalMsg
	}

	dbpool.UpdateRecord(staffNew)
	return httpdata.NewResponseSucc(&requestResponse_updateOrgStaff{
		Staff: staffNew,
	}), nil
}

// updateSelfStaffInfo 员工自己可以修改自己的信息
func (handler *requestHandler_updateOrgStaff) updateSelfStaffInfo(staff, staffExisting *orgtype.OrgStaff) (interface{}, httpdata.HttpError) {
	staffNew := staffExisting.AssignFields_self(staff)
	if staffNew == nil {
		return nil, httpdata.RcYouCanUpdateNickDeskOnly
	}

	dbpool.UpdateRecord(staffNew)

	response := httpdata.NewResponseSucc(&requestResponse_updateOrgStaff{
		Staff: staffNew,
	})
	return response, nil
}
