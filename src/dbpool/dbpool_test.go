package dbpool

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/lizisky/liziutils/utils"
// 	"lizisky.com/lizisky/src/basictypes/basictype"
// 	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
// 	eduorg "lizisky.com/lizisky/src/basictypes/orgtype_edu"
// 	"lizisky.com/lizisky/src/config"
// 	"lizisky.com/lizisky/src/dbpool/dbimpl"
// )

// func TestDB_impl(t *testing.T) {
// 	if !config.LoadConfig() {
// 		return
// 	}

// 	dbimpl.NewDBConnection(nil)
// 	defer dbimpl.DB.Close()

// 	// tmp_create_read_org()
// 	tmp_create_read_course()
// }

// func tmp_create_read_org() {
// 	org1 := &orgtype.Organization{
// 		Name:          "org+" + time.Now().String(),
// 		MajorCategory: 1,
// 		MajorName:     "major1234567890",
// 		// Images:        &basictype.ImageSlice{"images/abcd.png", "pic/efgh.png"},
// 		Images: &basictype.BasicSlice[string]{"img/abcd.png", "pic2/efgh.png"},
// 	}

// 	dbimpl.AddNewRecord(org1)

// 	fmt.Println("create - org1: ", utils.ToJSONIndent(org1))

// 	org2, _ := dbimpl.GetOrganizationByOID(org1.Oid)
// 	if org2 != nil {
// 		fmt.Println("read - org2: ", utils.ToJSONIndent(org2))
// 	}
// }

// func tmp_create_read_course() {

// 	// uintslice := []uint{1, 2, 3}
// 	// // fmt.Println("uintslice: ", utils.ToJSON(uintslice))

// 	// tmp, err := json.Marshal(uintslice)
// 	// fmt.Println("uintslice: ", string(tmp), err)

// 	// unslide := []uint{}
// 	// err = json.Unmarshal(tmp, &unslide)
// 	// fmt.Println("unslide: ", unslide, err)

// 	// return

// 	course1 := &eduorg.CourseInfo{
// 		Oid:       1,
// 		Name:      "course + " + time.Now().String(),
// 		Amount:    100,
// 		FirstDay:  1672532200000,
// 		StartTime: 54000000,
// 		Duration:  5000000,
// 		// Images:    &basictype.ImageSlice{"images/abcd.png", "pic/efgh.png"},
// 		// Images: &basictype.BasicSlice[string]{"img/abcd.png", "pic2/efgh.png"},
// 		// DaysInWeek: []uint8{1, 2, 3, 4},
// 		DaysInWeek: &basictype.BasicSlice[uint]{1, 2, 3},
// 		// DaysInWeek: &basictype.Uint8Slice{Days: []uint8{1, 2, 3, 4}},
// 		// DaysInWeek: []int{1, 2, 3, 4},
// 	}

// 	dbimpl.AddNewRecord(course1)
// 	fmt.Println("create - course1: ", utils.ToJSONIndent(course1))

// 	course2, _ := dbimpl.GetCourseInfo(course1.Cid)
// 	if course2 != nil {
// 		fmt.Println("read - course2-12: ", course2.DaysInWeek)
// 		fmt.Println("read - course2: ", utils.ToJSONIndent(course2))
// 	}

// }
