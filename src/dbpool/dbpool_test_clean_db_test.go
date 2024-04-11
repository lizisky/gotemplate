package dbpool

import (
	"fmt"
	"testing"

	"lizisky.com/lizisky/src/config"
	"lizisky.com/lizisky/src/dbpool/dbimpl"
)

func TestDB_impl_2(t *testing.T) {
	if !config.LoadConfig() {
		return
	}

	dbimpl.NewDBConnection(nil)
	// defer dbimpl.DB.Close()

	clean_db()
}

// func clean_db_2() {

// 	db := dbimpl.GetDB()

// 	var act []*accounts.Account
// 	db.Find(&act)
// 	db.Delete(&act)

// 	var org []*orgtype.Organization
// 	db.Find(&org)
// 	db.Delete(&org)

// 	var orgstaff []*orgtype.OrgStaff
// 	db.Find(&orgstaff)
// 	db.Delete(&orgstaff)

// 	var orgcourse []*eduorg.CourseInfo
// 	db.Find(&orgcourse)
// 	db.Delete(&orgcourse)

// 	var orgclass []*eduorg.ClassInfo
// 	db.Find(&orgclass)
// 	db.Delete(&orgclass)

// 	var orgaskleave []*eduorg.AskLeaveInfo
// 	db.Find(&orgaskleave)
// 	db.Delete(&orgaskleave)

// 	var orgcoursemember []*eduorg.CourseMembership
// 	db.Find(&orgcoursemember)
// 	db.Delete(&orgcoursemember)

// 	var orgcoursemembersignin []*eduorg.CourseMemberSignin
// 	db.Find(&orgcoursemembersignin)
// 	db.Delete(&orgcoursemembersignin)
// }

func clean_db() {
	// if !config.LoadConfig() {
	// 	return
	// }

	// dbimpl.NewDBConnection(nil)
	// defer dbimpl.DB.Close()

	tableNames, _ := listMySQLTables()
	if len(tableNames) == 0 {
		return
	}

	for _, tableName := range tableNames {
		// if tableName == "migrations" {
		// 	continue
		// }
		// fmt.Println(tableName)

		db := dbimpl.GetDB().Table(tableName).Delete(nil)
		fmt.Println("db: ", db.Error)
		// dbimpl.DB.Exec("TRUNCATE TABLE " + tableName)
	}

}

func listMySQLTables() ([]string, error) {

	var tableNames []string
	result := dbimpl.GetDB().Raw("SHOW TABLES").Scan(&tableNames)
	if result.Error != nil {
		return nil, result.Error
	}

	return tableNames, nil
}
