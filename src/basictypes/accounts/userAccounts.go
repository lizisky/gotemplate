// 这个文件主要定义个人账户相关的数据结构

package accounts

// Account Type for Register/Login/etc
const (
	AccountTypeBegin     = 0
	AccountTypeMobile    = 1
	AccountTypeWXOpenID  = 2
	AccountTypeWXUnionID = 3
	AccountTypeEmail     = 4
	AccountTypeAID       = 5
	AccountTypeEnd       = 6
)

// Account represent an account data
type Account struct {
	// Account ID
	Aid        uint64 `json:"aid,omitempty" gorm:"primary_key;auto_increment;unique;not null"`
	WxOpenID   string `json:"wxOpenID,omitempty" gorm:"unique" validate:"omitempty,min=20"` // WeChat Open ID
	Mobile     string `json:"mobile,omitempty" validate:"omitempty,numeric,len=11"`         // Mobile Number
	CreateDate int64  `json:"createDate,omitempty"`                                         // account create date
	Nickname   string `json:"nickName,omitempty" validate:"omitempty,min=1,max=16"`         // Nickname, ref: constants.MinLength_accountName
	Birthday   int64  `json:"birthday,omitempty"`                                           // birthday, maybe some's birthday before 1970.1.1
	Gender     uint8  `json:"gender,omitempty" validate:"oneof=0 1 2"`                      // 用户性别  0:unknown 1:male 2:female
	AvatarUrl  string `json:"avatarUrl,omitempty"`                                          // 头像
	Password   string `json:"pwd,omitempty"`                                                //
	Salt       string `json:"-"`                                                            // salt for password, Max Length: 32bytes
	Email      string `json:"email,omitempty" validate:"omitempty,email"`                   // email address

	// Name       string `json:"name,omitempty" validate:"omitempty,min=2,max=16"`             // Real name
	// Province   string `json:"province,omitempty"`                                           // 省份
	// City       string `json:"city,omitempty"`                                               // 市
	// District   string `json:"district,omitempty"`                                           // 县/区

	// IDcard   string `json:"idCard,omitempty" validate:"omitempty,len=18"` // ID Card,身份证号码
	// Country  string `json:"country,omitempty"`                            //
	// Language string `json:"language,omitempty"`                           // language
	// WxUnionID string `json:"wxUnionID,omitempty" validate:"omitempty,min=20"` // WeChat Union ID
}
