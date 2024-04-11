package httpdata

const (
	// code of success
	RC_OkCode = 1 // OK code

	// --------------------------------------------------------------------------------------------
	//
	// Error Code 分类：1
	// 这一类错误码，仅仅需要客户端把msg展示给用户，由用户执行某种操作，比如：输入的信息长度不对
	// 客户端不需要做其他的操作
	//
	RC_ShowErrMsg = 100 // 显示 msg 给用户
	//
	// Error Code 分类：1
	//

	// --------------------------------------------------------------------------------------------
	//
	// Error Code 分类：2
	// 这一类错误码，客户端发现 rc == RC_SystemErr，把 msg 写入日志即可，
	// 不需要客户端做其他特殊处理
	// 把这些错误码写入 msg，而不是直接在 RC 返回给客户端，
	// 除了 RC_SystemErr 以外，客户端不需要关心其他错误吗
	//
	RC_SystemErr = 201 // general system error
	//
	RC_SystemCrashAndRecovered      = 210 // 系统崩溃，然后自动重启
	RC_Write2DBFailed               = 211 // 写入数据库失败
	RC_SystemErrConveryResultFailed = 212 // system error for convert HTTP returned value to JSON failed
	RC_InvalidHttpRequestParams     = 213 // invalid request params
	//
	// Error Code 分类：2
	//

	// --------------------------------------------------------------------------------------------
	//
	// Error Code 分类：3
	// 错误码大于1000，需要客户端针对每一个错误码做出相应的动作
	//
	RC_NeedDeposit                  = 1001 // 需要用户充值，for example: 用户已经透支了，必须提醒用户
	RC_GetWechatOpenidError         = 1002 // 从微信服务器获取 wx open id 失败
	RC_SessionExpired               = 1103 // session expired
	RC_ParseHttpRequestParamsFailed = 1104 // convert request data failed, 这种错误是由于客户端上传到服务器端的数据格式不对所造成
	RC_ClientVersionExpired         = 1105 // 客户端版本已经过期了
	//
	// Error Code 分类：3
	//

	// code related with http request parameters
)

var (
	RcSyetemErr              = &HttpErrorData{RC_SystemErr, "系统错误:%d"}
	RcNotImplenment          = &HttpErrorData{RC_ShowErrMsg, "此功能尚未实现"}
	RcInvalidHttpRequestArgs = &HttpErrorData{RC_ShowErrMsg, "无效的请求参数"}
	RcClientVersionExpired   = &HttpErrorData{RC_ClientVersionExpired, "客户端版本已经已经不再被支持，请升级到最新版本"}

	RcConvertRequestParamDataErr = &HttpErrorData{RC_ParseHttpRequestParamsFailed, "请求参数转换错误"}
	RcSessionExpired             = &HttpErrorData{RC_SessionExpired, "您的登录凭证已经过期，请重新登录"}

	// code related with customer consume record in Org
	// RcYourBalanceIsZero = &HttpErrorData{RC_NeedDeposit, "您的余额为0了，请及时充值"}
	// RcYouHaveOverSpent  = &HttpErrorData{RC_NeedDeposit, "您已经透支了(%d)次，请及时充值"}

	RcNeednotUpdateIdenticalMsg = &HttpErrorData{RC_ShowErrMsg, "不需要更新相同的信息"}
	RcNoRleatedData             = &HttpErrorData{RC_ShowErrMsg, "没有相关数据"}

	// code related with general user operation
	RcAccountNotExisting     = &HttpErrorData{RC_ShowErrMsg, "用户(%v)不存在"}
	RcUserDidNotJoinAnyOrg   = &HttpErrorData{RC_ShowErrMsg, "用户(%d)还没有加入任何组织"}
	RcUserAlreadyMemberOfOrg = &HttpErrorData{RC_ShowErrMsg, "用户(%d)已经是组织(%d)的会员了"}
	RcUserIsNotStaffOfOrg    = &HttpErrorData{RC_ShowErrMsg, "组织(%d)中没有这个员工(%d)"}
	RcRejectUpdateRootOid    = &HttpErrorData{RC_ShowErrMsg, "不能更换顶层组织(%v)"}

	// code related with edu org
	RcUserOnlyUpdateSelfInfo = &HttpErrorData{RC_ShowErrMsg, "用户只能更新自己的信息"}
	RcGetPhoneNumberFailed   = &HttpErrorData{RC_ShowErrMsg, "获取手机号码失败，请稍后重试"}
	RcGetWechatOpenidError   = &HttpErrorData{RC_ShowErrMsg, "获取微信OpenID失败，请稍后重试"}

	// code related with general org operation
	RcOrgNotExisting          = &HttpErrorData{RC_ShowErrMsg, "组织(%d)不存在"}
	RcInvalidOrgParentAndRoot = &HttpErrorData{RC_ShowErrMsg, "无效的上级机构(%d,%d)"}
	RcInvalidOrgHierarchic    = &HttpErrorData{RC_ShowErrMsg, "无效的组织层级关系(%v,%v,%v)"}
	RcInvalidOrgCategory      = &HttpErrorData{RC_ShowErrMsg, "无效的机构分类(%d,%d)"}
	RcUserNoRightToDoThis     = &HttpErrorData{RC_ShowErrMsg, "用户(%d)没有权限在组织(%d)中执行这个动作"}

	// code related with org staffs
	RcUserAlreadyStaffOfOrg    = &HttpErrorData{RC_ShowErrMsg, "用户(%d)已经是组织(%d)的员工了"}
	RcYouCannotAddYourself     = &HttpErrorData{RC_ShowErrMsg, "不能添加自己为企业员工"}
	RcYouCannotRemoveYourself  = &HttpErrorData{RC_ShowErrMsg, "不能从企业员工中删除自己"}
	RcYouCanUpdateNickDeskOnly = &HttpErrorData{RC_ShowErrMsg, "你只能修改自己在企业中的昵称和说明信息"}

	// Course related error code
	// RcCourseNotExisting         = &HttpErrorData{RC_ShowErrMsg, "课程(%d)不存在"}
	// RcClassNotExisting          = &HttpErrorData{RC_ShowErrMsg, "没有这一节课(%d)的信息"}
	// RcUserIsNotMemberOfCousre   = &HttpErrorData{RC_ShowErrMsg, "用户(%d)还没有报名这个课程(%d)"}
	// RcCanOnlyChangeClassPeriod  = &HttpErrorData{RC_ShowErrMsg, "只能修改上课时间"}
	// RcClassPeriodConflict       = &HttpErrorData{RC_ShowErrMsg, "课时冲突,请检查上课时间，或者先删除有冲突的课时。另一节课的上课时间是(%v~%v)"}
	// RcNoEnrollInfo              = &HttpErrorData{RC_ShowErrMsg, "没有报名信息"}
	// RC_CannotDeleteClassStarted = &HttpErrorData{RC_ShowErrMsg, "不能删除已经开始或者结束的课程"}
	// RcExceedMaxClassAmount      = &HttpErrorData{RC_ShowErrMsg, "一个课程的最大课时数是(%d)"}
	// RcExceedWeekdays            = &HttpErrorData{RC_ShowErrMsg, "一周的时间只能是在0~6之间"}
	// RcMustInOneDay              = &HttpErrorData{RC_ShowErrMsg, "必须保证第一天的时间以及第一节课的开始时间和结束时间在同一天"}
	// RcNoRelatedClassses         = &HttpErrorData{RC_ShowErrMsg, "没有相关的上课信息"}
	// RcCancelSelfEnrollONLY      = &HttpErrorData{RC_ShowErrMsg, "只能取消自己的报名信息"}
	// RcRemainAmountNotZero       = &HttpErrorData{RC_ShowErrMsg, "剩余课时数不为0，不能取消课程报名"}
	// RcCourseOrgCannotBeChanged  = &HttpErrorData{RC_ShowErrMsg, "不能修改课程所属的组织"}

	// code related with leave info
	// RcNoLeaveInfo               = &HttpErrorData{RC_ShowErrMsg, "没有请假信息"}
	// RcOverAskLeaveScope         = &HttpErrorData{RC_ShowErrMsg, "这一节课(%d)不在您的请假范围内"}
	// RcYouHadAskLeaveInThisClass = &HttpErrorData{RC_ShowErrMsg, "这一节课(%d)你已经请过假了"}
	// RcExceedAskLeaveTime        = &HttpErrorData{RC_ShowErrMsg, "已经超过请假时间了。上课开始(%v)分钟前可以请假"}
	// RcNoRightToDeleteLeaveInfo  = &HttpErrorData{RC_ShowErrMsg, "您没有权限取消这个请假信息"}

	// code related with signin info
	// RcReadMemberSigninRecordFailed   = &HttpErrorData{RC_ShowErrMsg, "读取用户(%d)在课程(%d)中的签到信息失败"}
	// RcCanNotSigninFuture             = &HttpErrorData{RC_ShowErrMsg, "不能对未来的时间进行签到"}
	// RcCanNotSignin30DaysAgo          = &HttpErrorData{RC_ShowErrMsg, "不能对30天之前的时间签到"}
	// RcThisClassHasSigned             = &HttpErrorData{RC_ShowErrMsg, "你(%d)在这一节课(%d)已经签到过了"}
	// RcCannotUpdateSigninInfoYourself = &HttpErrorData{RC_ShowErrMsg, "不能更新自己的签到信息"}

	// RCorgNameLengthInvalid       = &HttpErrorData{RC_OrgNameLengthInvalid, "机构名字长度在4~64之间"}
	// RcNeednotUpdatAnything       = &HttpErrorData{RC_NeednotUpdateAnything, "不需要更新任何信息"}
)

// predefined error message for using easily
const (
	RcSystemErrConveryResultFailed_json_str = `{"rc":201,"msg":"系统错误:212"}` // RC_SystemErrConveryResultFailed
)
