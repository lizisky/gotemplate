package organization

import (
	"net/http"

	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/orgutil"
)

//
// Get Organization List
// parentOid == 0, 获取 Root Org List
// parentOid > 0, 获取某一个 parentOid 的 Sub Org List （只有第一下级，不包含更下一级的层级）
//
// Count:  Client希望获取 Org 的数量，最大只能获取到 50 个。 如果 count == 0 或 count > 50， 则最大返回 10 个

// MajorCategory: Org 主分类 ID
// SubCategory:   Org 子分类 ID
// 这两个参数如果等于 0，则表示不限制分类，返回所有分类的 Org List
// 如果大于 0， 则必须符合 getAllOrgCategories 接口返回的分类信息，否则返回错误。
//

// define http request param
type requestParam_getOrgList struct {
	Count         uint32 `json:"count,omitempty"`         // Client希望获取 Org 的数量，0 表示不限制
	ParentOid     uint64 `json:"parentOid,omitempty"`     // Parent Org OID
	MajorCategory uint32 `json:"majorCategory,omitempty"` // Org 主分类 ID
	SubCategory   uint32 `json:"subCategory,omitempty"`   // Org 子分类 ID
	// MajorCategory uint32 `json:"majorCategory,omitempty" validate:"gt=0,lt=100"`  // Org 主分类 ID
	// SubCategory   uint32 `json:"subCategory,omitempty" validate:"gt=100,lt=9999"` // Org 子分类 ID
}

// validate request parameters
func (param *requestParam_getOrgList) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

// define http request response
type requestResponse_getOrgList struct {
	OrgList []*orgtype.Organization `json:"orgList,omitempty"` // org list
}

// requestHandler_getOrgList implements the HttpHandler interface
type requestHandler_getOrgList struct {
}

func (handler *requestHandler_getOrgList) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_getOrgList) URL() string {
	return httpurl.Url_Org_Get_Org_List
}

func (handler *requestHandler_getOrgList) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_getOrgList
	return &param
}

func (handler *requestHandler_getOrgList) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_getOrgList)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	response := httpdata.NewResponseSucc(&requestResponse_getOrgList{
		OrgList: orgutil.GetOrgList(param.ParentOid, param.MajorCategory, param.SubCategory, param.Count, nil),
	})

	return response, nil
}
