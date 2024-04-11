// 这个文件主要定义组织账户相关的数据结构

package orgtype

import (
	"lizisky.com/lizisky/src/basictypes/basictype"
)

// Organization "一个组织"数据结构
type Organization struct {
	// Organization ID
	Oid uint64 `json:"oid,omitempty" gorm:"primary_key;auto_increment;unique;not null"`
	// parent Organization ID, 如果这个==0，说明这是一个 root org
	ParentOid uint64 `json:"parentOid,omitempty"`
	// root Organization ID, 如果这个==0，说明这 org 本身就是一个 root org
	// 如果 ParentOid > 0, 则 RootOid 必须 > 0，否则为无效数据
	RootOid uint64 `json:"rootOid,omitempty"`

	Name          string `json:"name,omitempty" validate:"omitempty,min=4,max=64"`      // org name
	MajorCategory uint32 `json:"majorCategory,omitempty"`                               // 一级分类ID, 不同领域可以自行定制
	MajorName     string `json:"majorName,omitempty" validate:"omitempty,min=2,max=64"` // 一级分类name
	SubCategory   uint32 `json:"subCategory,omitempty"`                                 // 二级分类ID, 不同领域可以自行定制
	SubName       string `json:"subName,omitempty" validate:"omitempty,min=2,max=64"`   // 二级分类name
	Logo          string `json:"logo,omitempty"`                                        // org logo URL
	Description   string `json:"description,omitempty" validate:"omitempty,min=4"`      // 组织介绍
	Address       string `json:"address,omitempty"`                                     // 地址
	PostCode      string `json:"postCode,omitempty"`                                    // 邮政编码
	Telephone     string `json:"telephone,omitempty"`                                   // 电话, desktop telphone, or, mobile telphone
	Email         string `json:"email,omitempty" validate:"omitempty,email"`            // 电子邮件
	Homepage      string `json:"homepage,omitempty"`                                    // 主页地址
	Qualification string `json:"qualification,omitempty"`                               // 资质
	BizCert       string `json:"bizCert,omitempty"`                                     // 经营许可证, image link
	BizLicence    string `json:"bizLicence,omitempty"`                                  // 营业执照图片文件名, image link
	BizPrimary    string `json:"bizPrimary,omitempty"`                                  // 主营业务
	CreateDate    int64  `json:"createDate,omitempty"`                                  // 创建日期

	// 这个Org的创建者的 Account ID
	Creator uint64 `json:"creator,omitempty"`
	// 这个Org的拥有者的 Account ID，
	// 初始创建的时候，Owner == Creator，后续 Owner 可以转让给其他人
	Owner       uint64  `json:"owner,omitempty"`
	State       uint8   `json:"state,omitempty"`                                    // 状态，未知/正常/等待审核/禁止/已经注销
	BankName    string  `json:"bankName,omitempty"`                                 // 开户银行 name
	BankAccount string  `json:"bankAccount,omitempty"`                              // 银行账号
	Longitude   float64 `json:"longitude,omitempty" validate:"omitempty,longitude"` // 经度
	Latitude    float64 `json:"latitude,omitempty" validate:"omitempty,latitude"`   // 纬度

	// org 相关图片, 这里是一个string数组，可以存储多张图片的 URL, 不可包含重复数据
	Images *basictype.ImageSlice `json:"images,omitempty" gorm:"type:json" validate:"omitempty,unique"`
	// Country string `json:"country,omitempty"` // 所属国家
}

// OrgStaff "组织的员工"数据结构
type OrgStaff struct {
	ID          uint64 `json:"-" gorm:"primary_key;auto_increment;unique;not null"`   // 暂时用不上
	Aid         uint64 `json:"aid,omitempty" validate:"gt=0"`                         // 员工 account id
	Aliasname   string `json:"aliasName,omitempty" validate:"omitempty,gte=1,lte=16"` // 员工在这个组织中的别名
	Mobile      string `json:"mobile,omitempty" validate:"omitempty,numeric,len=11"`  // Mobile Number
	Oid         uint64 `json:"oid,omitempty" validate:"gt=0"`                         // 所属组织的 org id
	ParentOid   uint64 `json:"parentOid,omitempty"`                                   // parent Organization ID, 如果这个==0，说明这个 Staff 直属于 root org
	RootOid     uint64 `json:"rootOid,omitempty"`                                     // root Organization ID, 如果这个==0，说明这个 Staff 直属于 oid 的 org
	Role        uint32 `json:"role,omitempty" validate:"omitempty,gte=1,lte=4"`       // 员工在这个Org中的角色
	RoleName    string `json:"roleName,omitempty" validate:"omitempty,gte=2,lte=16"`  // 员工在这个Org中的角色名，总经理、副总经理、主管、组长、助理，等等
	AddBy       uint64 `json:"addBy,omitempty"`                                       // 把这个员工加入这个Team或Org的Account ID
	AddByName   string `json:"AddByName,omitempty" validate:"omitempty,gte=2,lte=16"` // 把这个员工加入这个Team或Org的Account Name
	CreateDate  int64  `json:"createDate,omitempty"`                                  //
	Description string `json:"description,omitempty" validate:"omitempty,gte=4"`      // 说明信息
}
