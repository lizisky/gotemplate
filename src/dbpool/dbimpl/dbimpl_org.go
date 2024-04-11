package dbimpl

import (
	"errors"

	"gorm.io/gorm"
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
)

// GetOrganizationByOID get an organization info by ID
func GetOrganizationByOID(oid uint64) (*orgtype.Organization, error) {
	if oid == 0 {
		return nil, errors.New(invalidParameter)
	}

	/*
		var org *orgtype.Organization
		// fmt.Printf("111, poinger of org: %p\n", org)
		db := localDB.currDB.Where("oid = ?", oid).First(&org)
		// fmt.Printf("222, poinger of org: %p\n", org)

		// org := orgtype.Organization{}
		// fmt.Println("000, poinger of org: ", org)
		// fmt.Printf("111, poinger of org: %p\n", &org)
		// db := localDB.currDB.First(&org, oid)
		// fmt.Printf("222, poinger of org: %p\n", &org)

		if db.Error != nil {
			org = nil
		}
		return org, db.Error
	*/
	return get_org_record("oid = ?", oid)
}

func get_org_record(query interface{}, args ...interface{}) (*orgtype.Organization, error) {
	var org *orgtype.Organization
	db := localDB.currDB.Where(query, args...).First(&org)

	if db.Error != nil {
		org = nil
	}
	return org, db.Error
}

// IsOrgExisting checks is an organization existing or not
func IsOrgExisting(oid uint64) bool {
	if oid == 0 {
		return false
	}

	org := &orgtype.Organization{}
	db := localDB.currDB.Select("oid").First(org, oid)
	return db.Error == nil
}

// IsOrgExisting checks is an organization existing or not
func IsOrgExisting_withRootOid(oid, rootOid uint64) bool {
	if (oid == 0) || (rootOid == 0) {
		return false
	}

	if oid == rootOid {
		return IsOrgExisting(oid)
	}

	org := &orgtype.Organization{}
	db := localDB.currDB.Select("oid", "root_oid").First(org, oid)
	if db.Error != nil {
		return false
	}

	if org.RootOid != rootOid {
		return false
	}

	return true
}

// GetAllOrgList get an organization info by ID
// parentOid: 0 means get all root orgs
// parentOid: > 0 means get all sub orgs of parentOid
func GetOrgList(parentOid uint64, majorCategory, subCategory, count uint32, excludeOids []uint64) ([]*orgtype.Organization, error) {
	var db *gorm.DB

	// db := localDB.currDB.Where("parent_oid = ?", parentOid).Find(&list)
	if (majorCategory > 0) && (subCategory > 0) {
		if !orgtype.IsValidMajorCategory(majorCategory) || !orgtype.IsValidSubCategory(subCategory) {
			return nil, errors.New(invalidParameter)
		}
		db = db.Where("parent_oid = ? AND major_category = ? AND sub_category = ?", parentOid, majorCategory, subCategory)
	} else if majorCategory > 0 {
		if !orgtype.IsValidMajorCategory(majorCategory) {
			return nil, errors.New(invalidParameter)
		}
		db = db.Where("parent_oid = ? AND major_category = ?", parentOid, majorCategory)
	} else {
		db = localDB.currDB.Where("parent_oid = ?", parentOid)
	}

	if count < 1 || count > 51 {
		count = 10
	}
	if (excludeOids != nil) && (len(excludeOids) > 0) {
		db = db.Not(excludeOids)
	}

	var list []*orgtype.Organization
	db = db.Limit(int(count)).Find(&list)

	if db.Error != nil {
		list = nil
	}
	return list, db.Error
}

// GetOrgListByOwnerID 获取 Org Owner 是 owner_aid 的 OrgList
func GetOrgListByOwnerID(owner_aid uint64) ([]*orgtype.Organization, error) {
	if owner_aid == 0 {
		return nil, errors.New(invalidParameter)
	}

	var list []*orgtype.Organization
	db := localDB.currDB.Where("owner = ?", owner_aid).Find(&list)

	if db.Error != nil {
		list = nil
	}
	return list, db.Error
}

// GetOrgListByStaffID 获取 staff_aid 所在的 OrgList
func GetOrgListByStaffID(staff_aid uint64) ([]*orgtype.Organization, error) {
	if staff_aid == 0 {
		return nil, errors.New(invalidParameter)
	}

	var orgList []*orgtype.Organization
	db := localDB.currDB.Where("owner = ?", staff_aid).Find(&orgList)

	if db.Error != nil {
		orgList = nil
	}

	var staffList []*orgtype.OrgStaff
	db = localDB.currDB.Select("oid").Where("aid = ?", staff_aid).Find(&staffList)
	if (db.Error == nil) && (len(staffList) > 0) {
		for _, staff := range staffList {
			var org *orgtype.Organization
			db = localDB.currDB.Where("oid = ?", staff.Oid).First(&org)
			if (db.Error == nil) && (org != nil) {
				orgList = append(orgList, org)
			}
		}
	}

	return orgList, db.Error
}

// 判断一个用户(aid)是否已经是一个组织(oid)的 -- 工作人员
func IsStaffInOrg(aid, oid uint64) bool {
	if (aid == 0) || (oid == 0) {
		return false
	}

	var staff *orgtype.OrgStaff
	db := localDB.currDB.Select("aid").Where("aid = ? AND oid = ?", aid, oid).First(&staff)
	return db.Error == nil
}

// 读取一个组织(oid & parentOid)中的工作人员(aid)信息
func GetStaffInOrg(aid, oid uint64) (*orgtype.OrgStaff, error) {
	if (aid == 0) || (oid == 0) {
		return nil, errors.New(invalidParameter)
	}

	var staff *orgtype.OrgStaff
	db := localDB.currDB.Where("aid = ? AND oid = ?", aid, oid).First(&staff)
	if db.Error != nil {
		staff = nil
	}
	return staff, db.Error
}

// 从一个组织(oid & parentOid)中移除工作人员(aid)
func RemoveStaffFromOrg(aid, oid uint64) bool {
	if (aid == 0) || (oid == 0) {
		return false
	}

	db := localDB.currDB.Where("aid = ? AND oid = ?", aid, oid).Delete(&orgtype.OrgStaff{})
	return db.Error == nil
}

// get staff list info in Org
// @oid: Organization ID
// @parentOid: Parent-Organization ID. 如果parentOid为0，则是获取root org的stafflist
// @roles: 需要获取的在这个Org中的指定的角色，角色定义参见：src/basictypes/orgtype/org_common.go 如果没有传递 roles 参数，则获取oid&parentOid中所有staff list
func GetStaffListInOrg(oid, parentOid uint64, roles ...int) ([]*orgtype.OrgStaff, error) {
	if oid == 0 {
		return nil, errors.New(invalidParameter)
	}

	var staffList []*orgtype.OrgStaff

	queryClause := "oid = ? AND parent_oid = ?"
	db := localDB.currDB.Where(queryClause, oid, parentOid)

	if len(roles) > 0 {
		db = db.Where("role IN ?", roles)
	}

	db = db.Find(&staffList)

	if db.Error != nil {
		staffList = nil
	}
	return staffList, db.Error
}

// 获取一个员工在这个Org中的角色
// @aid: account ID
// @oid: Organization ID
// @roles: 需要获取的在这个Org中的指定的角色，如果没有传递 roles 参数，则获取oid&parentOid中所有staff list
// @return Role ID
// 角色定义参见：src/basictypes/orgtype/org_common.go
func GetStaffRoleInOrg(aid, oid uint64, roles ...uint32) uint32 {
	if (aid == 0) || (oid == 0) {
		return 0
	}

	var staff *orgtype.OrgStaff

	queryClause := "aid = ? AND oid = ?"
	db := localDB.currDB.Select("role").Where(queryClause, aid, oid)

	if len(roles) > 0 {
		db = db.Where("role IN ?", roles)
	}

	db = db.First(&staff)
	if db.Error != nil {
		return 0
	}
	return staff.Role
}
