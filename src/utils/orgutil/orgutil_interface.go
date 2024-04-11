package orgutil

import (
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/http/httpdata"
)

// IsUserHasRight check if user has right to do something
// @return nil if user has right, else return error
func IsUserHasRight(aid, oid uint64, roles ...uint32) httpdata.HttpError {
	return isUserHasRight_impl(aid, oid, roles...)
}

// ReadOrgInfo read org basic info
func ReadOrgInfo(oid uint64) *orgtype.Organization {
	return readOrgInfo_impl(oid)
}

// GetAllOrgList get an organization info by ID
// parentOid: 0 means get all root orgs
// parentOid: > 0 means get all sub orgs of parentOid
func GetOrgList(parentOid uint64, majorCategory, subCategory, count uint32, excludeOids []uint64) []*orgtype.Organization {
	return getOrgList_impl(parentOid, majorCategory, subCategory, count, excludeOids)
}

// BuildDownloadURL_org, build all images download URL for a organization
func BuildDownloadURL_org(org *orgtype.Organization) {
	buildDownloadURL_org_impl(org)
}

// --------------------------------------------------------------------------------------------------------
// checkOrgCategory 检查组织的分类信息是否正确
func CheckOrgCategory(majorCategory, subCategory uint32) bool {
	return checkOrgCategory_impl(majorCategory, subCategory)
}

// --------------------------------------------------------------------------------------------------------
// checkOrgStructure 检查组织的结构信息是否正确
func CheckOrgHierarchic(parentOid, rootOid uint64) bool {
	return checkOrgHierarchic_impl(parentOid, rootOid)
}

// --------------------------------------------------------------------------------------------------------
// checkOrgStructure 检查组织的结构信息是否正确
func CheckOrgHierarchic_full(oid, parentOid, rootOid uint64) bool {
	return CheckOrgHierarchic_full_impl(oid, parentOid, rootOid)
}
