package dbpool

import (
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/dbpool/dbimpl"
)

// ----------------------------------------------------------------------------
//
// DB operations for organization
//

// GetOrganizationByOID get an organization info by ID
func GetOrganizationByOID(oid uint64) (*orgtype.Organization, error) {
	return dbimpl.GetOrganizationByOID(oid)
}

// IsOrgExisting checks is an organization existing or not
func IsOrgExisting(oid uint64) bool {
	return dbimpl.IsOrgExisting(oid)
}

// IsOrgExisting checks is an organization existing or not
func IsOrgExisting_withRootOid(oid, rootOid uint64) bool {
	return dbimpl.IsOrgExisting_withRootOid(oid, rootOid)
}

// GetAllOrgList get an organization info by ID
// parentOid: 0 means get all root orgs
// parentOid: > 0 means get all sub orgs of parentOid
func GetOrgList(parentOid uint64, majorCategory, subCategory, count uint32, excludeOids []uint64) ([]*orgtype.Organization, error) {
	return dbimpl.GetOrgList(parentOid, majorCategory, subCategory, count, excludeOids)
}

// GetOrgListByOwnerID 获取 Org Owner 是 owner_aid 的 OrgList
func GetOrgListByOwnerID(owner_aid uint64) ([]*orgtype.Organization, error) {
	return dbimpl.GetOrgListByOwnerID(owner_aid)
}

// GetOrgListByStaffID
func GetOrgListByStaffID(staff_aid uint64) ([]*orgtype.Organization, error) {
	return dbimpl.GetOrgListByStaffID(staff_aid)
}

// 判断一个用户(aid)是否已经是一个组织(oid, parentOid)的 -- 工作人员
func IsStaffInOrg(aid, oid uint64) bool {
	return dbimpl.IsStaffInOrg(aid, oid)
}

// 读取一个组织(oid & parentOid)中的工作人员(aid)信息
func GetStaffInOrg(aid, oid uint64) (*orgtype.OrgStaff, error) {
	return dbimpl.GetStaffInOrg(aid, oid)
}

// 从一个组织(oid & parentOid)中移除工作人员(aid)
func RemoveStaffFromOrg(aid, oid uint64) bool {
	return dbimpl.RemoveStaffFromOrg(aid, oid)
}

// get staff list info in Org
// @oid: Organization ID
// @parentOid: Sub-Organization ID. 如果parentOid为0，则是获取root org的stafflist
// @roles: 需要获取的在这个Org中的指定的角色，如果没有传递 roles 参数，则获取oid&parentOid中所有staff list
// 角色定义参见：src/basictypes/orgtype/org_common.go
func GetStaffListInOrg(oid, parentOid uint64, roles ...int) ([]*orgtype.OrgStaff, error) {
	return dbimpl.GetStaffListInOrg(oid, parentOid, roles...)
}

// 获取一个员工在这个Org中的角色
// @aid: account ID
// @oid: Organization ID
// @roles: 需要获取的在这个Org中的指定的角色，如果没有传递 roles 参数，则获取oid&parentOid中所有staff list
// @return Role ID
// 角色定义参见：src/basictypes/orgtype/org_common.go
func GetStaffRoleInOrg(aid, oid uint64, roles ...uint32) uint32 {
	return dbimpl.GetStaffRoleInOrg(aid, oid, roles...)
}
