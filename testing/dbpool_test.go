package main

// import (
// 	"fmt"
// 	"math"
// 	"testing"
// 	"time"
//
// 	"github.com/lizisky/liziutils/utils"
// 	"lizisky.com/lizisky/src/basictypes/accounts"
// 	"lizisky.com/lizisky/src/basictypes/basictype"
// 	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
// 	eduorg "lizisky.com/lizisky/src/basictypes/orgtype_edu"
// 	"lizisky.com/lizisky/src/config"
// 	"lizisky.com/lizisky/src/dbpool"
// 	"lizisky.com/lizisky/src/dbpool/dbimpl"
// 	"lizisky.com/lizisky/src/utils/orgutil"
// )
//
// func TestDB_impl(t *testing.T) {
// 	if !config.LoadConfig() {
// 		return
// 	}
//
// 	dbimpl.NewDBConnection(nil)
// 	defer dbimpl.DB.Close()
//
// 	// try_TestDB_add_query_account(t)
// 	// try_TestDB_add_query_org(t)
//
// 	// try_user_apply_toJoinOrg(t)
// 	// try_get_org_membership_list(t)
// 	// try_get_org_membership_of_account(t)
//
// 	// try_get_all_org(t)
//
// 	// try_validator()
//
// 	// tmp_course_info()
//
// 	try_org_staffs(t)
// }
//
// func try_validator() {
//
// 	account := accounts.Account{
// 		Aid:      1,
// 		WxOpenID: ";alksjdf;laksjd;lfajs;dlfj;alsdf",
// 		Name:     "aaklajsdklfhalsj",
// 	}
//
// 	fmt.Println(utils.IsValidStruct(account))
// }
//
// func try_TestDB_add_query_account(t *testing.T) {
//
// 	// fmt.Println("generate uuid:", utils.NewUUID(), utils.NewRandomUUID())
//
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	act1 := &accounts.Account{
// 		WxOpenID:  utils.NewRandomUUID(),
// 		WxUnionID: utils.GenerateRandomStr(nil),
// 	}
//
// 	dbpool.AddNewRecord(act1)
// 	// fmt.Println("testing 1 ---", utils.ToJSONIndent(act1))
//
// 	fmt.Printf("Is account existing, %d, %v\n", act1.Aid+1, dbpool.IsAccountExisting(act1.Aid+1))
// 	fmt.Printf("Is account existing, %d, %v\n", act1.Aid, dbpool.IsAccountExisting(act1.Aid))
//
// 	// act2, err := GetAccountByWXopenid(act1.WxOpenID)
// 	// if err == nil {
// 	// 	fmt.Println("----------------------- testing 21 ---\n", utils.ToJSONIndent(act2))
// 	// } else {
// 	// 	fmt.Println("----------------------- testing 22 ---", err)
// 	// }
//
// 	act3, err := dbpool.GetAccountByAID(act1.Aid)
// 	if err == nil {
// 		fmt.Println("----------------------- testing 31 ---\n", utils.ToJSONIndent(act3))
// 	} else {
// 		fmt.Println("----------------------- testing 32 ---", err)
// 	}
// }
//
// func try_TestDB_add_query_org(t *testing.T) {
//
// 	// fmt.Println("generate uuid:", utils.NewUUID(), utils.NewRandomUUID())
//
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	org1 := &orgtype.Organization{
// 		// Oid: 2,
// 		Name:        utils.NewRandomUUID(),
// 		Description: utils.GenerateRandomStr(nil),
// 	}
//
// 	err := dbpool.AddNewRecord(org1)
// 	// err := SaveRecord(org1)
// 	fmt.Println("testing 1 ---", utils.ToJSONIndent(org1), err)
//
// 	fmt.Printf("Is org existing, %d, %v\n", org1.Oid+1, dbpool.IsOrgExisting(org1.Oid+1))
// 	fmt.Printf("Is org existing, %d, %v\n", org1.Oid, dbpool.IsOrgExisting(org1.Oid))
//
// 	org22, err := dbpool.GetOrganizationByOID(org1.Oid)
// 	if err == nil {
// 		fmt.Println("----------------------- testing 21 ---", utils.ToJSONIndent(org22))
// 	} else {
// 		fmt.Println("----------------------- testing 22 ---", err)
// 	}
// }
//
// func try_get_all_org(t *testing.T) {
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	list, err := dbpool.GetOrgList(0)
// 	fmt.Println(err)
//
// 	for _, org := range list {
// 		fmt.Println(utils.ToJSONIndent(org))
// 	}
// }
//
// func try_user_apply_toJoinOrg(t *testing.T) {
//
// 	// fmt.Println("generate uuid:", utils.NewUUID(), utils.NewRandomUUID())
//
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	aid := uint64(5)
// 	oid := uint64(4)
// 	membership := &eduorg.CourseMembership{
// 		Aid:         aid,
// 		Oid:         oid,
// 		CreateDate:  time.Now().UnixMilli(),
// 		Nickname:    "亚历山大",
// 		Description: "亚历山大今天假日喜丽体育",
// 	}
//
// 	err := dbpool.AddNewRecord(membership)
// 	fmt.Println("testing 1 ---", err)
// 	fmt.Println("testing 2 ---", utils.ToJSONIndent(membership))
//
// 	fmt.Printf("Is member %d in org, %d --, %v\n", aid+10, oid, dbpool.IsMemberInOrg(aid+10, oid))
// 	fmt.Printf("Is member %d in org, %d --, %v\n", aid, oid, dbpool.IsMemberInOrg(aid, oid))
// }
//
// func try_get_org_membership_list(t *testing.T) {
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	oid := uint64(2)
// 	list, _ := dbpool.GetMembershipListInOrg(oid)
// 	fmt.Println("try_get_org_membership_list ---, member count:", len(list))
// 	fmt.Println(utils.ToJSONIndent(list))
// }
//
// func try_get_org_membership_of_account(t *testing.T) {
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	aid := uint64(5)
// 	list, _ := dbpool.GetOrgListByMemberID(aid)
// 	fmt.Println("try_get_org_membership_of_account ---, member count:", len(list))
// 	fmt.Println(utils.ToJSONIndent(list))
// }
//
// func tmp_course_info() {
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// defer dbimpl.DB.Close()
//
// 	cid := uint64(1)
// 	imgs := &basictype.BasicSlice[string]{"c123", utils.NewRandomUUID(), "c123"}
// 	days := &basictype.BasicSlice[uint8]{1, 2, 3}
//
// 	course := &eduorg.CourseInfo{
// 		Cid:         cid,
// 		Oid:         1,
// 		Name:        "abcde",
// 		CreateDate:  1,
// 		Description: "descritpion 12312312",
// 		Images:      imgs,
// 		DaysInWeek:  days,
// 		Amount:      1,
// 	}
//
// 	dbimpl.AddNewRecord(course)
//
// 	// c2, _ := dbimpl.GetCourseInfo(cid, 1, 0, 1)
// 	c2, _ := dbimpl.GetCourseInfo(cid)
// 	fmt.Println("read data from db:\n", utils.ToJSONIndent(c2))
//
// 	ant := uint32(0xFFFFFFFF)
// 	fmt.Printf("%X, %X, %v\n", ant, math.MaxUint32, ant == math.MaxUint32)
//
// 	abcd := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
// 	abcd15 := time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC)
// 	fmt.Println(abcd.UnixMilli(), abcd15.UnixMilli()-abcd.UnixMilli())
//
// 	fmt.Println("has duplicate 1", days.HadDuplicate())
// 	days = &basictype.BasicSlice[uint8]{1, 2, 3, 2}
// 	fmt.Println("has duplicate 2", days.HadDuplicate())
// 	fmt.Println("has duplicate 3", utils.IsValidStruct(course))
// }
//
// // func buildDB_Config() *config.Configuration {
// // 	return &config.Configuration{
// // 		ServerAddr: ":8081",
// // 		MySQL: config.DBConfig{
// // 			DBHost:     "localhost:3306",
// // 			DBUser:     "root",
// // 			DBUserPwd:  "123456",
// // 			DBDatabase: "lizisky001",
// // 			MaxConn:    10,
// // 		},
// // 	}
// // }
//
// func try_org_staffs(t *testing.T) {
//
// 	// dbimpl.NewDBConnection(buildDB_Config())
// 	// role := dbpool.GetStaffRoleInOrg(2, 1, 0, 2, 3, 4, 5)
// 	// fmt.Println("GetStaffRoleInOrg result is:", role)
//
// 	// list, _ := dbpool.GetStaffListInOrg(1, 0, 2, 4)
// 	// fmt.Println("GetStaffListInOrg result is:", utils.ToJSONIndent(list))
//
// 	info, _ := orgutil.ReadOrgDeepInfo(1, 2)
// 	fmt.Println("ReadOrgDeepInfo result is:", utils.ToJSONIndent(info))
// }
