package dbimpl

import (
	"errors"

	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/basictypes/constants"
)

// get account info by WX Open ID
func GetAccountByWXopenid(wxOpenID string) (*accounts.Account, error) {
	if len(wxOpenID) < constants.MinLength_wxOpenID {
		return nil, errors.New(invalidParameter)
	}
	return get_account_record("wx_open_id = ?", wxOpenID)
}

// get account info by mobile number
func GetAccountByMobile(mobile string) (*accounts.Account, error) {
	if len(mobile) < constants.LengthMobileNumber {
		return nil, errors.New(invalidParameter)
	}
	return get_account_record("mobile = ?", mobile)
}

// get account info by account id
func GetAccountByAID(aid uint64) (*accounts.Account, error) {
	if aid < 1 {
		return nil, errors.New(invalidParameter)
	}
	return get_account_record("aid = ?", aid)
}

// generic function to get account info
func get_account_record(query interface{}, args ...interface{}) (*accounts.Account, error) {
	var act *accounts.Account
	db := localDB.currDB.Where(query, args...).First(&act)

	if db.Error != nil {
		act = nil
	}
	return act, db.Error
}

// IsAccountExisting checks is an organization existing or not
func IsAccountExisting(aid uint64) bool {
	if aid < 1 {
		return false
	}
	var act *accounts.Account
	db := localDB.currDB.Model(act).Select("aid").First(&act, aid)
	return db.Error == nil
	// var act *accounts.Account
	// db := localDB.currDB.Model(act).Select("aid").First(&act, aid)
	// return db.Error == nil
}
