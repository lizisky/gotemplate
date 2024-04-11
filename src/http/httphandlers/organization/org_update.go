package organization

import (
	"net/http"

	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/orgutil"
)

//
// Update Organization Information
//

// define http request param
type requestParam_updateOrgInfo struct {
	// SessionToken string               `json:"sessionToken,omitempty"`                 // session token
	OrgInfo orgtype.Organization `json:"orginfo,omitempty" validate:"omitempty"` // org info
}

// validate request parameter
func (param *requestParam_updateOrgInfo) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	if param.OrgInfo.Oid == 0 {
		return httpdata.RcOrgNotExisting.WithArgs(0)
	}
	return nil
}

// define http request response
type requestResponse_updateOrgInfo struct {
	OrgInfo *orgtype.Organization `json:"orginfo,omitempty" validate:"omitempty"` // org info
}

// requestHandler_updateOrgInfo implements the HttpHandler interface
type requestHandler_updateOrgInfo struct {
}

func (handler *requestHandler_updateOrgInfo) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_updateOrgInfo) URL() string {
	return httpurl.Url_Org_Update_Organization
}

func (handler *requestHandler_updateOrgInfo) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_updateOrgInfo
	return &param
}

// Handle 处理请求
func (handler *requestHandler_updateOrgInfo) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_updateOrgInfo)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return nil, httpdata.RcSessionExpired
	// }

	return handler.updateOrg(rdata.Magic.SessionInfo.Aid, &param.OrgInfo)
}

// --------------------------------------------------------------------------------------------------------
// updateOrg 更新一个原有组织的信息
func (handler *requestHandler_updateOrgInfo) updateOrg(aid uint64, orginfo *orgtype.Organization) (interface{}, httpdata.HttpError) {
	if (orginfo.MajorCategory > 0) || (orginfo.SubCategory > 0) {
		// 因为这里是 update orginfo，所以不设置这两个数据表示不更新这两个数据
		// 但是如果有一个设置了，就必须检查其有效性
		// check category ID, MUST set major category ID
		if !orgutil.CheckOrgCategory(orginfo.MajorCategory, orginfo.SubCategory) {
			return nil, httpdata.RcInvalidOrgCategory.WithArgs(orginfo.MajorCategory, orginfo.SubCategory)
		}
	}

	existingOrgInfo, _ := dbpool.GetOrganizationByOID(orginfo.Oid)
	if existingOrgInfo == nil {
		// DB 中没有 orgInfo
		return nil, httpdata.RcOrgNotExisting.WithArgs(orginfo.Oid)
	}
	// existingOrgInfo read from DB based on orginfo.Oid, so it's safe to use existingOrgInfo.Oid

	// 不允许更换root org
	if orginfo.RootOid != existingOrgInfo.RootOid {
		return nil, httpdata.RcRejectUpdateRootOid.WithArgs(orginfo.RootOid)
	}
	// 如果隶属关系发生变更，检查新的隶属关系是否有效
	if orginfo.ParentOid != existingOrgInfo.ParentOid {
		if !orgutil.CheckOrgHierarchic(orginfo.ParentOid, orginfo.RootOid) {
			return nil, httpdata.RcInvalidOrgParentAndRoot.WithArgs(orginfo.ParentOid, orginfo.RootOid)
		}
		// if (orginfo.ParentOid == 0) && (orginfo.RootOid == 0) 视作把 orginfo 从一个 sub org 提升为 root org
	}

	// 只有 Owner/ViceOwner 有权限更新组织信息
	err2 := orgutil.IsUserHasRight(aid, existingOrgInfo.Oid, orgtype.RoleOwner, orgtype.RoleViceOwner)
	if err2 != nil {
		return nil, err2
	}

	response := httpdata.NewResponseSucc()
	// 从 DB 中获得了原有的 orgInfo
	orginfoNew := existingOrgInfo.AssignFields(orginfo)
	if orginfoNew == nil {
		// 信息完全相同，没必要 update, do nothing
		return nil, httpdata.RcNeednotUpdateIdenticalMsg
	}

	// update an existing organization information
	// 有部分信息不相同，do update
	dbpool.UpdateRecord(orginfoNew)

	response.Data = &requestResponse_updateOrgInfo{
		OrgInfo: orginfoNew,
	}
	return response, nil
}
