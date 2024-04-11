package accountUtil

import (
	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/dbpool"
)

// GetAccountByWXopenid returns the account information by the given wxOpenID
func GetAccountByWXopenid(wxOpenID string) *accounts.Account {
	account, _ := dbpool.GetAccountByWXopenid(wxOpenID)
	return post_process_account(account)
}

// GetAccountByMobile returns the account information by the given mobile
func GetAccountByMobile(mobile string) *accounts.Account {
	account, _ := dbpool.GetAccountByMobile(mobile)
	return post_process_account(account)
}

// GetAccountByAID returns the account information by the given aid
func GetAccountByAID(aid uint64) *accounts.Account {
	account, _ := dbpool.GetAccountByAID(aid)
	return post_process_account(account)
}
