package commonHandlers

import (
	"net/http"

	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
)

// 读取一个 account 的 homepage data
//
// Mode: 0 invalid parameter
// Mode: 1 get homepage data for a staff
// Mode: 2 Get homepage data for a member
// Mode: 3 include mode 1 and mode 2

// define http request param
type requestParam_readHomepageData struct {
	Aid  uint64 `json:"aid,omitempty" validate:"gt=0"` //  account id
	Mode int    `json:"mode,omitempty" validate:"gt=0"`
}

// validate request parameters
func (param *requestParam_readHomepageData) IsValid() httpdata.HttpError {
	return nil
}

// define http request response
type requestResponse_readHomepageData struct {
	// data for org staff
	// StaffData *orgmix.Homepage_data_for_staff `json:"staffData,omitempty"`

	// // data for org membership
	// MembershipData *orgmix.Homepage_data_for_membership `json:"membershipData,omitempty"`
}

// requestHandler_readHomepageData implements the HttpHandler interface
type requestHandler_readHomepageData struct {
}

func (handler *requestHandler_readHomepageData) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_readHomepageData) URL() string {
	return httpurl.Url_Common_read_homepage_data
}

func (handler *requestHandler_readHomepageData) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_readHomepageData
	return &param
}

// Handle : provide implementation of interface HttpHandler's Handle() method
func (handler *requestHandler_readHomepageData) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	// param, ok := rdata.Data.(*requestParam_readHomepageData)
	// if !ok {
	// 	return nil, httpdata.RcConvertRequestParamDataErr
	// }

	// list, _ := orgDeepUtil.readHomepageDataList(param.Oid, param.RemainAmount, param.DeepDegree)
	// if list == nil {
	// 	return nil, httpdata.RcNoEnrollInfo
	// }
	// act := accountUtil.GetAccountByAID(param.Aid)
	// is_new_account := (act == nil) || !act.HasMboile()

	response := &requestResponse_readHomepageData{}

	// if param.Mode&0x01 != 0 {
	// 	response.StaffData = orgDeepUtil.ReadHomepageData_For_Staff(param.Aid)
	// }

	// if param.Mode&0x02 != 0 {
	// 	response.MembershipData = orgDeepUtil.ReadHomepageData_For_Membership(param.Aid, is_new_account)
	// }

	return httpdata.NewResponseSucc(response), nil
}
