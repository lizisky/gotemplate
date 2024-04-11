package httpurl

const (
	// general api version prefix
	apiver = "/v1"

	// for general debug usage, echo message
	EchoMessage = apiver + "/demo/echo"

	//
	// API for common features
	//
	apiCommon = apiver + "/cmn"
	// pull bootup information: /v1/cmn/pull/bootup/info
	Url_Common_Pull_Bootup_Info = apiCommon + "/pull/bootup/info"
	// get qiniu upload token: /v1/cmn/get/qiniu/upload/token
	Url_Common_get_qiniu_upload_token = apiCommon + "/get/qiniu/upload/token"
	// pull bootup information: /v1/cmn/read/homepage/data
	Url_Common_read_homepage_data = apiCommon + "/read/homepage/data"

	//
	// API for Wechat related features
	//
	apiWechat = apiver + "/wx"
	// update WeChat auth code: /v1/wx/login/auth/code
	Url_Wechat_Login_Auth_Code = apiWechat + "/login/auth/code"
	// get wx phone number: /v1/wx/get/phone/number
	Url_Wechat_GetPhoneNumber = apiWechat + "/get/phone/number"
	// update mobile number: /v1/wx/upload/mobile/number
	// Url_Wechat_UploadMobileNumber = apiWechat + "/upload/mobile/number"

	//
	// API for general user accounts
	//
	apiAccount = apiver + "/act"
	// search user info: /v1/act/search/user
	Url_Account_Search_User = apiAccount + "/search/user"
	// update user info: /v1/act/update/user/info
	Url_Account_Update_User_Info = apiAccount + "/update/user/info"

	//
	// API for Organization operation
	//
	apiPetHouse = apiver + "/pet"
	// get all org categories: /v1/org/get/all/org/categories
	// Url_Org_Get_Org_Categories = apiOrg + "/get/all/org/categories"
	// create organization info: /v1/org/create/org
	Url_Create_pet_house = apiPetHouse + "/create/pethouse"
	// read organization info: /v1/org/read/orginfo
	// Url_Org_Read_Organization = apiOrg + "/read/orginfo"
	// update organization info: /v1/org/update/orginfo
	// Url_Org_Update_Organization = apiOrg + "/update/orginfo"

	//
	// API for Organization operation
	//
	apiOrg = apiver + "/org"
	// get all org categories: /v1/org/get/all/org/categories
	Url_Org_Get_Org_Categories = apiOrg + "/get/all/org/categories"
	// create organization info: /v1/org/create/org
	Url_Org_Create_Organization = apiOrg + "/create/org"
	// read organization info: /v1/org/read/orginfo
	Url_Org_Read_Organization = apiOrg + "/read/orginfo"
	// update organization info: /v1/org/update/orginfo
	Url_Org_Update_Organization = apiOrg + "/update/orginfo"

	// add staff in org: /v1/org/add/staff/in/org
	Url_Org_add_staff_in_org = apiOrg + "/add/staff/in/org"
	// read staff in org: /v1/org/read/staff/in/org
	Url_Org_read_staff_in_org = apiOrg + "/read/staff/in/org"
	// update staff in org: /v1/org/update/staff/in/org
	Url_Org_update_staff_in_org = apiOrg + "/update/staff/in/org"
	// remove staff in org: /v1/org/remove/staff/in/org
	Url_Org_remove_staff_in_org = apiOrg + "/remove/staff/in/org"

	// user apply to join org: /v1/org/apply/to/join/org
	// Url_Org_Apply_to_Join_Org = apiOrg + "/apply/to/join/org"
	// get membership list: /v1/org/get/membership/list
	Url_Org_Get_Membership_List = apiOrg + "/get/membership/list"
	// get org list: /v1/org/get/org/list
	Url_Org_Get_Org_List = apiOrg + "/get/org/list"

	//
	// API for Organization Of edu/training type
	//
	apiEduOrg = apiver + "/org/edu"
	// create course in org: /v1/org/edu/create/course
	Url_Edu_Org_Create_Course = apiEduOrg + "/create/course"
	// read course in org: /v1/org/edu/read/course
	Url_Edu_Org_Read_Course = apiEduOrg + "/read/course"
	// create course in org: /v1/org/edu/update/course
	Url_Edu_Org_Update_Course = apiEduOrg + "/update/course"

	// read class info: /v1/org/edu/read/class
	Url_Edu_Org_read_class = apiEduOrg + "/read/class"
	// update class info: /v1/org/edu/update/class
	Url_Edu_Org_update_classinfo = apiEduOrg + "/update/class"
	// delete Class: /v1/org/edu/delete/class
	Url_Edu_Org_delete_class = apiEduOrg + "/delete/class"
	// get coming classes list: /v1/org/edu/get/coming/classes
	Url_Edu_Org_Get_Coming_Classes = apiEduOrg + "/get/coming/classes"
	// get classes in course: /v1/org/edu/get/classes/in/course
	Url_Edu_Org_Get_Classes_in_course = apiEduOrg + "/get/classes/in/course"
	// get passed classes in course: /v1/org/edu/get/passed/classes/in/course
	Url_Edu_Org_Get_Passed_Classes_in_course = apiEduOrg + "/get/passed/classes/in/course"

	// create ask for leave: /v1/org/edu/create/leaveinfo
	Url_Edu_Org_create_leaveinfo = apiEduOrg + "/create/leaveinfo"
	// read leave: /v1/org/edu/read/leaveinfo
	Url_Edu_Org_read_leaveinfo = apiEduOrg + "/read/leaveinfo"
	// delete leave: /v1/org/edu/delete/leaveinfo
	Url_Edu_Org_delete_leave = apiEduOrg + "/delete/leaveinfo"

	// member enroll course: /v1/org/edu/member/enroll/course
	Url_Edu_Org_member_enroll_course = apiEduOrg + "/member/enroll/course"
	// update course membership info: /v1/org/edu/update/course/member/info
	Url_Org_Update_Course_MembershipInfo = apiEduOrg + "/update/course/member/info"
	// read membership info in course: /v1/org/edu/read/membership/in/course
	Url_Edu_Org_read_membership_in_course = apiEduOrg + "/read/membership/in/course"
	// delete membership info in course: /v1/org/edu/delete/membership/in/course
	Url_Edu_Org_delete_membership_in_course = apiEduOrg + "/delete/membership/in/course"
	// course membership sign in course: /v1/org/edu/member/signin
	Url_Edu_Org_Course_Member_Signin = apiEduOrg + "/member/signin"
	// course membership sign in course: /v1/org/edu/member/signin/update
	Url_Edu_Org_Course_Member_Signin_Update = apiEduOrg + "/member/signin/update"
	// get member sign in record: /v1/org/edu/get/member/signin/record
	Url_Org_Get_Member_Signin_Record = apiEduOrg + "/get/member/signin/record"
	// read membership info in org: /v1/org/edu/read/membership/in/org
	Url_Edu_Org_read_membership_in_org = apiEduOrg + "/read/membership/in/org"

	// create course in org: /v1/org/edu/class/closed/create
	Url_Edu_Org_ClassClosed_Create = apiEduOrg + "/class/closed/create"
)
