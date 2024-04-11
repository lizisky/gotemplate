package organization

import (
	"net/http"

	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
)

//
// Get All Organization Categories
//

// define http request param
type requestParam_getAllOrgCategories struct {
	// SessionToken string `json:"sessionToken,omitempty"` // session token
}

// validate request parameters
func (param *requestParam_getAllOrgCategories) IsValid() httpdata.HttpError {
	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return httpdata.RcSessionExpired
	// }
	return nil
}

// define http request response
type requestResponse_getAllOrgCategories struct {
	OrgCategories orgtype.OrgMajorCategory `json:"orgCategories,omitempty"` // org list
}

// requestHandler_getAllOrgCategories implements the HttpHandler interface
type requestHandler_getAllOrgCategories struct {
}

func (handler *requestHandler_getAllOrgCategories) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_getAllOrgCategories) URL() string {
	return httpurl.Url_Org_Get_Org_Categories
}

func (handler *requestHandler_getAllOrgCategories) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_getAllOrgCategories
	return &param
}

func (handler *requestHandler_getAllOrgCategories) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	response := httpdata.NewResponseSucc(requestResponse_getAllOrgCategories{
		OrgCategories: orgtype.GetOrgCategories(),
	})

	return response, nil
}
