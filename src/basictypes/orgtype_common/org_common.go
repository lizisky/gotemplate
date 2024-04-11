package orgtype

// 这个文件主要定义 org 中一些通用的定义

// 定义一个组织中各个成员的角色
const (
	RoleUnknown    = 0 // 未知角色
	RoleOwner      = 1 // Owner
	RoleViceOwner  = 2 // Vice Owner
	RoleAssistant  = 3 // 助理、秘书等
	RoleTeamMember = 4 // general team member, 一般工作人员
	RoleMembership = 5 // membership, 即：会员/客户
	RoleEnd        = 6 // 边界值
)

// 状态定义,
// 适用于 Organization.State，CourseInfo.State
// 添加新的value之后，必须调整StateEnd，确保StateEnd是最大值
const (
	StateUnknown        = 0 // 未知状态
	StateNormal         = 1 // 正常
	StatePaused         = 2 // 暂停
	StateClosed         = 3 // 结束
	StateRecruiting     = 4 // 正在发展新会员
	StateRecruitClosed  = 5 // 停止发展新会员
	StateWaitingApprove = 6 // 等待审批
	StateEnd            = 7 // 边界值
)
