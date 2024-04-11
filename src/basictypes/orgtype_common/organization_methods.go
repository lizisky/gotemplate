// 这个文件主要定义组织账户相关的数据结构

package orgtype

import (
	"github.com/lizisky/liziutils/utils"
	"lizisky.com/lizisky/src/basictypes/constants"
)

const (
	minLengthOfOrgField = 2
)

func (org *Organization) AssignFields(assign *Organization) *Organization {
	if (assign == nil) || (org.Oid != assign.Oid) || (assign.IsValid() != nil) {
		return nil
	}

	changed := false
	// if org.ParentOid == assign.ParentOid, 赋值一次没有任何影响
	// if assign.ParentOid == 0, 等同于认为，这个org被提升为 root org
	// if org.ParentOid != assign.ParentOid, 等同于认为，这个org在组织结构中，被划分到其他部门去了
	if org.ParentOid != assign.ParentOid {
		org.ParentOid = assign.ParentOid
		changed = true
	}

	// 不允许更换 RootOid
	// if org.RootOid != assign.RootOid {
	// 	org.RootOid = assign.RootOid
	// 	changed = true
	// }

	// 每一个string的有效性都必须做判断，因为 assign.IsValid() 中，一个 name == “” 也认为是有效的
	if (len(assign.Name) > minLengthOfOrgField) && (org.Name != assign.Name) {
		org.Name = assign.Name
		changed = true
	}
	if (assign.MajorCategory > 0) && (org.MajorCategory != assign.MajorCategory) {
		org.MajorCategory = assign.MajorCategory
		changed = true
	}
	if (len(assign.MajorName) > minLengthOfOrgField) && (org.MajorName != assign.MajorName) {
		org.MajorName = assign.MajorName
		changed = true
	}
	if (assign.SubCategory > 0) && (org.SubCategory != assign.SubCategory) {
		org.SubCategory = assign.SubCategory
		changed = true
	}
	if (len(assign.SubName) > minLengthOfOrgField) && (org.SubName != assign.SubName) {
		org.SubName = assign.SubName
		changed = true
	}
	// if (len(assign.Country) > minLengthOfOrgField) && (org.Country != assign.Country) {
	// 	org.Country = assign.Country
	// 	changed = true
	// }
	if (len(assign.Logo) > minLengthOfOrgField) && (org.Logo != assign.Logo) {
		org.Logo = assign.Logo
		changed = true
	}
	if (len(assign.Description) > minLengthOfOrgField) && (org.Description != assign.Description) {
		org.Description = assign.Description
		changed = true
	}
	if (len(assign.Address) > minLengthOfOrgField) && (org.Address != assign.Address) {
		org.Address = assign.Address
		changed = true
	}
	if (len(assign.PostCode) > minLengthOfOrgField) && (org.PostCode != assign.PostCode) {
		org.PostCode = assign.PostCode
		changed = true
	}
	if (len(assign.Telephone) > minLengthOfOrgField) && (org.Telephone != assign.Telephone) {
		org.Telephone = assign.Telephone
		changed = true
	}
	if (len(assign.Email) > minLengthOfOrgField) && (org.Email != assign.Email) {
		org.Email = assign.Email
		changed = true
	}
	if (len(assign.Homepage) > minLengthOfOrgField) && (org.Homepage != assign.Homepage) {
		org.Homepage = assign.Homepage
		changed = true
	}
	if (len(assign.Qualification) > minLengthOfOrgField) && (org.Qualification != assign.Qualification) {
		org.Qualification = assign.Qualification
		changed = true
	}
	if (len(assign.BizCert) > minLengthOfOrgField) && (org.BizCert != assign.BizCert) {
		org.BizCert = assign.BizCert
		changed = true
	}
	if (len(assign.BizLicence) > minLengthOfOrgField) && (org.BizLicence != assign.BizLicence) {
		org.BizLicence = assign.BizLicence
		changed = true
	}
	if (len(assign.BizPrimary) > minLengthOfOrgField) && (org.BizPrimary != assign.BizPrimary) {
		org.BizPrimary = assign.BizPrimary
		changed = true
	}
	if (assign.Owner > 0) && (org.Owner != assign.Owner) {
		org.Owner = assign.Owner
		changed = true
	}
	if (assign.State > 0) && (org.State != assign.State) {
		org.State = assign.State
		changed = true
	}
	if (len(assign.BankName) > minLengthOfOrgField) && (org.BankName != assign.BankName) {
		org.BankName = assign.BankName
		changed = true
	}
	if (len(assign.BankAccount) > minLengthOfOrgField) && (org.BankAccount != assign.BankAccount) {
		org.BankAccount = assign.BankAccount
		changed = true
	}
	if (assign.Longitude > float64(0)) && (org.Longitude != assign.Longitude) {
		org.Longitude = assign.Longitude
		changed = true
	}
	if (assign.Latitude > float64(0)) && (org.Latitude != assign.Latitude) {
		org.Latitude = assign.Latitude
		changed = true
	}

	if (assign.Images != nil) && assign.Images.IsValid() {
		org.Images = assign.Images
		changed = true
	}

	if changed {
		return org
	}

	return nil
}

// IsValid 判断是否有效
func (org *Organization) IsValid() error {
	return utils.IsValidStruct(org)
}

// IsOwner 判断是否是机构的所有者
func (org *Organization) IsOwner(aid uint64) bool {
	return org.Owner == aid
}

// update staff info, assign fields
func (staff *OrgStaff) AssignFields(assign *OrgStaff) *OrgStaff {
	if (assign == nil) || (staff.Oid != assign.Oid) || (staff.Aid != assign.Aid) {
		return nil
	}

	changed := false
	if (len(assign.Aliasname) > constants.MinLength_accountName) && (staff.Aliasname != assign.Aliasname) {
		staff.Aliasname = assign.Aliasname
		changed = true
	}

	if (len(assign.Mobile) > constants.LengthMobileNumber) && (staff.Mobile != assign.Mobile) {
		staff.Mobile = assign.Mobile
		changed = true
	}

	if staff.Role != assign.Role {
		staff.Role = assign.Role
		changed = true
	}

	if (len(assign.RoleName) > constants.MinLength_accountName) && (staff.RoleName != assign.RoleName) {
		staff.RoleName = assign.RoleName
		changed = true
	}

	if (len(assign.Description) > constants.MinLength_description) && (staff.Description != assign.Description) {
		staff.Description = assign.Description
		changed = true
	}

	if changed {
		return staff
	}

	return nil
}

// 企业员工修改自己的信息
func (staff *OrgStaff) AssignFields_self(assign *OrgStaff) *OrgStaff {
	if (assign == nil) || (staff.Oid != assign.Oid) || (staff.Aid != assign.Aid) {
		return nil
	}

	changed := false
	if (len(assign.Aliasname) > constants.MinLength_accountName) && (staff.Aliasname != assign.Aliasname) {
		staff.Aliasname = assign.Aliasname
		changed = true
	}

	if (len(assign.Description) > constants.MinLength_description) && (staff.Description != assign.Description) {
		staff.Description = assign.Description
		changed = true
	}

	if changed {
		return staff
	}

	return nil
}
