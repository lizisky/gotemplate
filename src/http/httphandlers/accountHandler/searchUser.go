package accountHandler

import (
	"net/http"

	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/accountUtil"
)

//
// search 用户信息
// 可以按照两种模式 search 用户信息：
// 1. 通过 AID search 用户信息
// 2. 通过手机号码 search 用户信息
//

const (
	search_user_by_aid    = 1
	search_user_by_mobile = 2
)

// define http request param
type requestParam_searchUser struct {
	SearchMode int    `json:"searchMode,omitempty" validate:"oneof=1 2"`            // 1: search user by AID  2: search user by Mobile
	Aid        uint64 `json:"aid,omitempty"`                                        // Account ID
	Mobile     string `json:"mobile,omitempty" validate:"omitempty,numeric,len=11"` // Mobile Number
}

// validate request parameters
func (param *requestParam_searchUser) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

type requestResponse_searchUser struct {
	Account *accounts.Account `json:"account,omitempty"` //
}

// requestHandler_searchUser implements the HttpHandler interface
type requestHandler_searchUser struct {
}

func (handler *requestHandler_searchUser) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_searchUser) URL() string {
	return httpurl.Url_Account_Search_User
}

func (handler *requestHandler_searchUser) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_searchUser
	return &param
}

func (handler *requestHandler_searchUser) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_searchUser)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	return handler.searchUser(param.SearchMode, param.Aid, param.Mobile)
}

// read user information
func (handler *requestHandler_searchUser) searchUser(mode int, aid uint64, mobile string) (interface{}, httpdata.HttpError) {
	var accountRead *accounts.Account
	var info interface{}

	switch mode {
	case search_user_by_aid:
		accountRead = accountUtil.GetAccountByAID(aid)
		info = aid
	case search_user_by_mobile:
		accountRead = accountUtil.GetAccountByMobile(mobile)
		info = mobile
	}

	if accountRead == nil {
		return nil, httpdata.RcAccountNotExisting.WithArgs(info)
	}

	response := httpdata.NewResponseSucc(&requestResponse_searchUser{
		Account: accountRead,
	})
	return response, nil
}
