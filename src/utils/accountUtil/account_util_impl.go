package accountUtil

import (
	"lizisky.com/lizisky/src/basictypes/accounts"
	"lizisky.com/lizisky/src/utils/qiniuUtil"
)

// post_process_account post-process the account information
func post_process_account(account *accounts.Account) *accounts.Account {
	if account == nil {
		return nil
	}

	if len(account.AvatarUrl) > 0 {
		account.AvatarUrl = qiniuUtil.BuildDownloadURL(account.AvatarUrl)
	}
	return account
}
