package orgtype

// define organization major category item
type OrgMajorCategoryItem struct {
	// major category name
	Name string `json:"name"`
	// sub category list, key: sub category id, value: sub category name
	SubCategory map[uint32]string `json:"subCategory"`
}

// define organization major category
type OrgMajorCategory map[uint32]OrgMajorCategoryItem

// define org major category
// major category ID scope: 1 ~ 100
const (
	OrgMajorCategory_Begin              = 0
	OrgMajorCategory_EducationTraining  = 1 // 教育培训
	OrgMajorCategory_BeautyHairdressing = 2 // 美容美发
	OrgMajorCategory_AutomobileService  = 3 // 汽车服务
	OrgMajorCategory_Housekeeping       = 4 // 家政服务
	// OrgMajorCategory_HouseholdAppliance  = 5 // 家电维修
	// OrgMajorCategory_HouseholdRepair     = 6 // 家居维修
	// OrgMajorCategory_HouseholdMoving     = 7 // 家居搬运
	// OrgMajorCategory_HouseholdCleaning   = 8 // 家居保洁
	// OrgMajorCategory_HouseholdDecoration = 9 // 家居装修
	OrgMajorCategory_End = 100 // end
)

// define org sub category
// sub category ID scope: start from: 101
const (
	// OrgMajorCategory_EducationTraining  = 1 // 教育培训
	OrgSubCategory_Begin                     = 100
	OrgSubCategory_EducationTraining_Sport   = 101 // 体育（足球/网球/乒乓球/游泳/羽毛球/等）
	OrgSubCategory_EducationTraining_Art     = 102 // 艺术（声乐/乐器/美术/等）
	OrgSubCategory_EducationTraining_English = 103 // 英语（四六级/口语/职场/出国/等）
	OrgSubCategory_EducationTraining_Culture = 104 // 文化课（语文/数学/英语/物理/化学/地理/历史/生物/等）
	OrgSubCategory_EducationTraining_Other   = 199 // 其他

	// OrgMajorCategory_BeautyHairdressing = 2 // 美容美发
	OrgSubCategory_BeautyHairdressing_Beauty       = 201 // 美容
	OrgSubCategory_BeautyHairdressing_Hairdressing = 202 // 美发
	OrgSubCategory_BeautyHairdressing_Other        = 250 // 其他

	// OrgMajorCategory_AutomobileService  = 3 // 汽车服务
	OrgSubCategory_AutomobileService_Wash   = 251 // 洗车
	OrgSubCategory_AutomobileService_Repair = 252 // 维修
	OrgSubCategory_AutomobileService_Other  = 300 // 其他

	// OrgMajorCategory_Housekeeping      = 4 // 家政服务
	OrgSubCategory_End = 9999
)

// IsValidMajorCategory returns true if majorCategory is valid
func IsValidMajorCategory(majorCategory uint32) bool {
	return (majorCategory > OrgMajorCategory_Begin && majorCategory < OrgMajorCategory_End)
}

// IsValidSubCategory returns true if subCategory is valid
func IsValidSubCategory(subCategory uint32) bool {
	return (subCategory > OrgSubCategory_Begin && subCategory < OrgSubCategory_End)
}

// GetOrgCategories returns the list of organization categories
// 所有 MajorCategory ID 线性编码
// 所有 SubCategory ID 线性编码
func GetOrgCategories() OrgMajorCategory {
	orgCategories := make(OrgMajorCategory, 10)

	// MajorCategory: 1 "教育培训"
	// SubCategory: ID (1 ~ 50)
	orgCategories[OrgMajorCategory_EducationTraining] = OrgMajorCategoryItem{
		Name: "教育培训",
		SubCategory: map[uint32]string{
			OrgSubCategory_EducationTraining_Sport:   "体育（足球/网球/乒乓球/游泳/羽毛球/等）",
			OrgSubCategory_EducationTraining_Art:     "艺术（声乐/乐器/美术/等）",
			OrgSubCategory_EducationTraining_English: "英语（四六级/口语/职场/出国/等）",
			OrgSubCategory_EducationTraining_Culture: "文化课（语文/数学/英语/物理/化学/地理/历史/生物/等）",
			OrgSubCategory_EducationTraining_Other:   "其他",
		},
	}

	// MajorCategory: 2 "美容美发"
	// SubCategory: ID (51 ~ 100)
	orgCategories[OrgMajorCategory_BeautyHairdressing] = OrgMajorCategoryItem{
		Name: "美容美发",
		SubCategory: map[uint32]string{
			OrgSubCategory_BeautyHairdressing_Beauty:       "美容",
			OrgSubCategory_BeautyHairdressing_Hairdressing: "美发",
			OrgSubCategory_BeautyHairdressing_Other:        "其他",
		},
	}

	// MajorCategory: 3 "汽车服务"
	// SubCategory: ID (101 ~ 150)
	orgCategories[OrgMajorCategory_AutomobileService] = OrgMajorCategoryItem{
		Name: "汽车服务",
		SubCategory: map[uint32]string{
			OrgSubCategory_AutomobileService_Wash:   "洗车",
			OrgSubCategory_AutomobileService_Repair: "保养",
			OrgSubCategory_AutomobileService_Other:  "其他",
		},
	}
	return orgCategories
}
