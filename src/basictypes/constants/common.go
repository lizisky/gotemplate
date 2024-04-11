package constants

// HTTP Client Type definition
const (
	ClientTypeUnkown            = 0 // unknown type
	ClientTypeWechatMiniAndroid = 1 // Wechat Mini Program on Android
	ClientTypeWechatMiniIOS     = 2 // Wechat Mini Program on iOS
	ClientTypeWechatMini        = 3 // Wechat Mini Program on other platform
	ClientTypeAndroid           = 4 // Android App
	ClientTypeIOS               = 5 // iOS App
	ClientTypeWeb               = 6 // Web Page
	ClientTypeH5                = 7 // H5 Page on Mobile Device
	ClientTypeEnd               = 8 // boundary end of client type
)

// 登录账号类型
const (
	LoginTypeInvalid      = 0 // invalid login type
	LoginTypeAID          = 1 // Account ID
	LoginTypeMobile       = 2 // Mobile
	LoginTypeSessionToken = 3 // Account ID
	LoginTypeEMail        = 4 // Email
	AccountTypeWXOpenID   = 5
	AccountTypeWXUnionID  = 6
	LoginTypeEnd          = 7 //  boundary end of login type
)

// general constant definition
const (
	// length of mobile number
	LengthMobileNumber = 11
	// length of password, it's original password hashed
	LengthPassword = 32
	// length of general signature
	LengthSignature = 32
)

const (
	// length of token-hash in HTTP request header
	// LEN_HEADER_TOKEN_HASH = 64
	// length of signature in HTTP request header
	// LEN_HEADER_SIGNATURE = 64

	// length of sessio token, that is a UUID without '-' symbol
	// Length_of_SessionToken = 32

	// org name length
	MinLength_orgName = 4
	MaxLength_orgName = 64

	// account nickName length
	MinLength_accountName = 1
	MaxLength_accountName = 16

	// length of description
	MinLength_description = 4

	// min length of avatar
	MinLength_avatar_url = 10
)

// const related with WeChat
const (
	// note: wechat官方没有给出这些数据的长度定义，这里都是根据经验值来定义的

	// minimal length of WeChat Open ID
	MinLength_wxOpenID = 20
	// minimal length of WeChat Union ID
	MinLength_wxUnionID = 20
	// minimal length of WeChat Login Auth Code
	MinLength_wxLoginAuthCode = 10
)
