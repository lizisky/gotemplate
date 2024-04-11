package orgutil

import (
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/utils/qiniuUtil"
)

// IsUserHasRight check if user has right to do something
func isUserHasRight_impl(aid, oid uint64, roles ...uint32) httpdata.HttpError {
	if !dbpool.IsAccountExisting(aid) {
		return httpdata.RcAccountNotExisting.WithArgs(aid)
	}

	// 不需要再次判断Oid > 0，因为前面环节的 IsValid 已经判断过了
	org, err := dbpool.GetOrganizationByOID(oid)
	if (err != nil) || (org == nil) {
		// DB 中没有这个 org
		return httpdata.RcOrgNotExisting.WithArgs(oid)
	}
	if org.IsOwner(aid) {
		return nil
	}

	// 当前用户不是这个Org的Owner，继续判断这个用户是不是这个Org的其他Owner/ViceOwner
	role := dbpool.GetStaffRoleInOrg(aid, oid, roles...)
	// if (orgtype.RoleOwner != role) && (role != orgtype.RoleViceOwner) {
	// 	return httpdata.RcUserNoRightToDoThis.WithArgs(aid, oid)
	// }
	if (role <= orgtype.RoleUnknown) || (role >= orgtype.RoleEnd) {
		// role 不在合法范围内
		return httpdata.RcUserNoRightToDoThis.WithArgs(aid, oid)
	}
	if len(roles) > 0 {
		// 循环遍历roles，如果role在roles中，返回success
		for _, r := range roles {
			if role == r {
				return nil
			}
		}
	}
	return httpdata.RcUserNoRightToDoThis.WithArgs(aid, oid)
}

// ReadOrgInfo read org basic info
func readOrgInfo_impl(oid uint64) *orgtype.Organization {
	orgInfo, _ := dbpool.GetOrganizationByOID(oid)
	if orgInfo == nil {
		return nil
	}

	buildDownloadURL_org_impl(orgInfo)

	return orgInfo
}

// GetAllOrgList get an organization info by ID
// parentOid: 0 means get all root orgs
// parentOid: > 0 means get all sub orgs of parentOid
func getOrgList_impl(parentOid uint64, majorCategory, subCategory, count uint32, excludeOids []uint64) []*orgtype.Organization {
	orglist, _ := dbpool.GetOrgList(parentOid, majorCategory, subCategory, count, excludeOids)
	if len(orglist) > 0 {
		for _, org := range orglist {
			buildDownloadURL_org_impl(org)
		}
	}
	return orglist
}

// buildDownloadURL_org_impl, build all images download URL for a course
func buildDownloadURL_org_impl(org *orgtype.Organization) {
	qiniuUtil.BuildDownloadURL_forImgSlice(org.Images)
	if len(org.Logo) > 0 {
		org.Logo = qiniuUtil.BuildDownloadURL(org.Logo)
	}

}

// --------------------------------------------------------------------------------------------------------
// checkOrgCategory 检查组织的分类信息是否正确
func checkOrgCategory_impl(majorCategory, subCategory uint32) bool {
	orgCategories := orgtype.GetOrgCategories()
	major, ok := orgCategories[majorCategory]
	if ok {
		// 可以没有 sub category ID
		if subCategory > 0 {
			_, ok = major.SubCategory[subCategory]
		}
		return ok
	}

	return false
}

// --------------------------------------------------------------------------------------------------------
// checkOrgStructure 检查组织的结构信息是否正确
// 这里仅检查 ParentOid 和 RootOid 的关系是否正确
func checkOrgHierarchic_impl(parentOid, rootOid uint64) bool {
	// 顶级组织，ParentOid 和 RootOid 必须为 0
	if (parentOid == 0) && (rootOid == 0) {
		return true
	}

	// 现在，ParentOid 和 RootOid 至少有一个大于 0，
	// 实际上，两者必须同时大于 0，否则，层级关系错误
	if (parentOid == 0) || (rootOid == 0) {
		return false
	}

	// 两者相等的情况，说明这个org是 rootOid 的第一级下级组织
	// 这种情况，只要 rootOid org 存在，即可
	if parentOid == rootOid {
		return dbpool.IsOrgExisting(rootOid)
	}

	// 两者大 > 0，且不相等的情况，说明这个 org 是 rootOid 的第二级或更深的下级组织
	// 这种情况，只要验证 rootOid 和 parentOid 的关系即可
	return dbpool.IsOrgExisting_withRootOid(parentOid, rootOid)
}

// --------------------------------------------------------------------------------------------------------
// checkOrgStructure 检查组织的结构信息是否正确
// 这里检查 Oid, ParentOid 和 RootOid 三者的关系是否正确
func CheckOrgHierarchic_full_impl(oid, parentOid, rootOid uint64) bool {
	if oid == 0 {
		return false
	}
	if !checkOrgHierarchic_impl(parentOid, rootOid) {
		return false
	}

	org, _ := dbpool.GetOrganizationByOID(oid)
	if org != nil {
		return (org.ParentOid == parentOid) && (org.RootOid == rootOid)
	}

	return false
}
