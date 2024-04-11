package wxutil

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/glog"
	"lizisky.com/lizisky/src/config"
	"lizisky.com/lizisky/src/utils/httpclient"
)

const (
	format_get_code       = "https://api.weixin.qq.com/sns/jscode2session?grant_type=authorization_code&appid=%s&secret=%s&js_code=%s"
	format_getAccessToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	format_getPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
)

// GetAccessToken
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func getAccessToken_impl() (string, error) {
	cfg := config.GetConfig()
	aurl := fmt.Sprintf(format_getAccessToken, cfg.WXMini.AppID, cfg.WXMini.Secret)
	body, err := httpclient.DoHttpGet(aurl)

	if (err == nil) && (len(body) > 0) {
		type wxTokenInfo struct {
			AccessToken string `json:"access_token"`
			Expires     int    `json:"expires_in"`
		}

		var sinfo wxTokenInfo
		err = json.Unmarshal(body, &sinfo)
		if err == nil {
			// fmt.Println("----------- GetAccessToken", utils.ToJSONIndent(sinfo))
			return sinfo.AccessToken, nil
		}
	}
	return "", err
}

// getPhoneNumber_impl
func getPhoneNumber_impl(codeOfGetPhoneNum string) (string, error) {
	accessToken, err := getAccessToken_impl()
	if err != nil {
		return "", err
	}

	phoneNumber, err := getPhoneNumber_impl_with_accessToken(accessToken, codeOfGetPhoneNum)
	return phoneNumber, err
}

// GetPhoneNumber
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func getPhoneNumber_impl_with_accessToken(accessToken, codeOfGetPhoneNum string) (string, error) {
	type wxGetPhoneNumberRequest struct {
		Code string `json:"code"`
	}

	type wxGetPhoneNumberWaterMark struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	}

	type wxGetPhoneNumberInfo struct {
		PhoneNumber     string                    `json:"phoneNumber"`
		PurePhoneNumber string                    `json:"purePhoneNumber"`
		CountryCode     string                    `json:"countryCode"`
		WaterMark       wxGetPhoneNumberWaterMark `json:"watermark"`
	}

	type wxGetPhoneNumberResponse struct {
		ErrCode   int                  `json:"errcode"`
		ErrMsg    string               `json:"errmsg"`
		PhoneInfo wxGetPhoneNumberInfo `json:"phone_info"`
	}

	code := &wxGetPhoneNumberRequest{Code: codeOfGetPhoneNum}
	bodyData, _ := json.Marshal(code)

	aurl := fmt.Sprintf(format_getPhoneNumber, accessToken)
	body, err := httpclient.DoHttpPost(aurl, bodyData)

	if (err == nil) && (len(body) > 0) {
		var sinfo wxGetPhoneNumberResponse
		err = json.Unmarshal(body, &sinfo)
		// fmt.Println("----------- GetPhoneNumber", utils.ToJSONIndent(sinfo))
		if err == nil {
			if sinfo.ErrCode == 0 {
				return sinfo.PhoneInfo.PhoneNumber, nil
			}
			err = fmt.Errorf("error code:%d, info:%s", sinfo.ErrCode, sinfo.ErrMsg)
		}
	}
	return "", err
}

// GetWXopenID
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
func getWXopenID_impl(authCode string) (session_key, openid string, err error) {
	cfg := config.GetConfig()
	aurl := fmt.Sprintf(format_get_code, cfg.WXMini.AppID, cfg.WXMini.Secret, authCode)
	body, err := httpclient.DoHttpGet(aurl)

	if (err == nil) && (len(body) > 0) {
		type wxSessionInfo struct {
			Session_key string `json:"session_key"`
			Openid      string `json:"openid"`
			Unionid     string `json:"unionid"`
			ErrCode     int    `json:"errcode"`
			ErrMsg      string `json:"errmsg"`
		}

		var sinfo wxSessionInfo
		err = json.Unmarshal(body, &sinfo)
		if (err == nil) && (sinfo.ErrCode == 0) {
			return sinfo.Session_key, sinfo.Openid, nil
		} else {
			glog.Info("get wx openid failed:", err)
			if err == nil {
				err = errors.New(sinfo.ErrMsg)
			}
			// fmt.Println("get wx openid failed:", err)
		}
	}
	return "", "", err
}
