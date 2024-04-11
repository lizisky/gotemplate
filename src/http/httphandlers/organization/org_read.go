package organization

import (
	"net/http"

	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/orgutil"
)

//
// Get Organization detailed info 获取一个机构的详细信息
//
// deep degree == 0: 只获取当前机构的基本信息
// deep degree > 0: 希望获取的组织机构的深度，最大可获取的org层级是9层.
// 包括：机构基本信息，员工列表，客户/会员列表，开设的业务类别信息列表（对于教育培训类机构，是开设的课程信息），下级org列表等
// 最大可获取的org层级是9层，数据结构则形成了一颗树，类似于真实的一个企业的组织结构
//

// define http request param
type requestParam_readOrgInfo struct {
	Oid uint64 `json:"oid,omitempty" validate:"gt=0"` // org id

	// deep degree, 希望获取的组织机构的深度，最大可获取的org层级是9层
	// 0: 只获取当前机构的基本信息
	DeepDegree int `json:"deepDegree,omitempty" validate:"gte=0,lt=10"`
}

// validate request parameters
func (param *requestParam_readOrgInfo) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

// define http request response
// type requestResponse_readOrgInfo struct {
// 	OrgDetailedInfo *orgtype.OrgDeepInfo `json:"org,omitempty"` // org list
// }

// requestHandler_readOrgInfo implements the HttpHandler interface
type requestHandler_readOrgInfo struct {
}

func (handler *requestHandler_readOrgInfo) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_readOrgInfo) URL() string {
	return httpurl.Url_Org_Read_Organization
}

func (handler *requestHandler_readOrgInfo) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_readOrgInfo
	return &param
}

func (handler *requestHandler_readOrgInfo) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_readOrgInfo)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// 仅读取组织机构信息
	if param.DeepDegree < 1 {
		return handler.readOrgOnly(param.Oid)
	}

	// 读取组织机构的详细信息
	// orginfo, _ := orgDeepUtil.ReadOrgDeepInfo(param.Oid, param.DeepDegree)
	// if orginfo == nil {
	// 	return nil, httpdata.RcOrgNotExisting.WithArgs(param.Oid)
	// }

	// return httpdata.NewResponseSucc(orginfo), nil
	return nil, httpdata.RcOrgNotExisting.WithArgs(param.Oid)
}

// readOrgOnly 仅读取组织机构信息
func (handler *requestHandler_readOrgInfo) readOrgOnly(oid uint64) (interface{}, httpdata.HttpError) {
	orginfo := orgutil.ReadOrgInfo(oid)
	if orginfo == nil {
		return nil, httpdata.RcOrgNotExisting.WithArgs(oid)
	}

	// response := httpdata.NewResponseSucc(&orgmix.OrgDeepInfo{
	// 	OrgInfo: orginfo,
	// })

	// return response, nil
	return nil, httpdata.RcOrgNotExisting.WithArgs(oid)
}
