package wxutil

//
// GetAccessToken
//
func GetAccessToken() (string, error) {
	return getAccessToken_impl()
}

//
// GetPhoneNumber
//
func GetPhoneNumber(codeOfGetPhoneNum string) (string, error) {
	return getPhoneNumber_impl(codeOfGetPhoneNum)
}

//
// GetWXopenID
//
func GetWXopenID(authCode string) (session_key, openid string, err error) {
	return getWXopenID_impl(authCode)
}
