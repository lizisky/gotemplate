package dbpool

import (
	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/dbpool/dbimpl"
)

// ----------------------------------------------------------------------------
//
// DB operations for user account
//

func GetAccountByWXopenid(wxOpenID string) (*accounts.Account, error) {
	return dbimpl.GetAccountByWXopenid(wxOpenID)
}

func GetAccountByMobile(mobile string) (*accounts.Account, error) {
	return dbimpl.GetAccountByMobile(mobile)
}

func GetAccountByAID(aid uint64) (*accounts.Account, error) {
	return dbimpl.GetAccountByAID(aid)
}

// IsAccountExisting checks is an organization existing or not
func IsAccountExisting(aid uint64) bool {
	return dbimpl.IsAccountExisting(aid)
}
