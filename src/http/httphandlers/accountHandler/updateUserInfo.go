package accountHandler

import (
	"net/http"

	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
	"lizisky.com/lizisky/src/utils/qiniuUtil"
)

//
// update user information
//

// define http request param
type requestParam_updateUserInfo struct {
	AccountInfo accounts.Account `json:"account,omitempty"` //
}

// validate request parameters
func (param *requestParam_updateUserInfo) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_ShowErrMsg, err)
	// }
	return nil
}

type requestResponse_updateUserInfo struct {
	Account *accounts.Account `json:"account,omitempty"` // account info
}

// requestHandler_updateUserInfo implements the HttpHandler interface
type requestHandler_updateUserInfo struct {
}

func (handler *requestHandler_updateUserInfo) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_updateUserInfo) URL() string {
	return httpurl.Url_Account_Update_User_Info
}

func (handler *requestHandler_updateUserInfo) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_updateUserInfo
	return &param
}

// Handle http request
func (handler *requestHandler_updateUserInfo) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	param, ok := rdata.Data.(*requestParam_updateUserInfo)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	return handler.updateUserInfo(rdata.Magic.SessionInfo, param)
}

// update user information
func (handler *requestHandler_updateUserInfo) updateUserInfo(sinfo *sessiondata.SessionInfo, param *requestParam_updateUserInfo) (interface{}, httpdata.HttpError) {
	if sinfo.Aid != param.AccountInfo.Aid {
		// 用户只能更新自己的信息
		return nil, httpdata.RcUserOnlyUpdateSelfInfo
	}

	account, _ := dbpool.GetAccountByAID(sinfo.Aid)
	if account == nil {
		// DB 中没有这个 account
		return nil, httpdata.RcAccountNotExisting.WithArgs(sinfo.Aid)
	}

	accountSave := account.AssignFields(&param.AccountInfo)
	if accountSave == nil {
		// 信息完全相同，没必要 update, do nothing
		return nil, httpdata.RcNeednotUpdateIdenticalMsg
	}

	dbpool.UpdateRecord(accountSave)

	// update Nickname in Session
	if sinfo.Nickname != param.AccountInfo.Nickname {
		sinfo.Nickname = param.AccountInfo.Nickname
		// sessionmgr.AddSession(param.SessionToken, sinfo) // update Nickname in Session
	}

	if len(accountSave.AvatarUrl) > 0 {
		accountSave.AvatarUrl = qiniuUtil.BuildDownloadURL(accountSave.AvatarUrl)
	}

	response := httpdata.NewResponseSucc(&requestResponse_updateUserInfo{
		Account: accountSave,
	})
	return response, nil
}
