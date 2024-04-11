// 这个文件主要定义个人账户相关的数据结构

package accounts

import (
	"github.com/go-playground/validator/v10"
	"lizisky.com/lizisky/src/basictypes/constants"
)

// RemoveSensitiveInfo remove sensitive information in the data
// func (act *Account) RemoveSensitiveInfo() {
// act.Password = ""
// act.Salt = ""
// act.IDcard = ""
// act.Email = ""
// act.Avatarurl = ""
// act.WxOpenID = ""
// act.WxUnionID = ""
// act.Mobile = ""
// }

// 如果所有这些关键数据都没有，则认为是 empty account
func (act *Account) IsEmpty() bool {
	return (act.Aid < 1) &&
		(len(act.WxOpenID) < constants.MinLength_wxOpenID) &&
		(len(act.Mobile) < 1) //&&
	// (len(act.Email) < 1) &&
	// (len(act.IDcard) < 1)
	// (len(act.WxUnionID) < constants.MinLength_wxUnionID) &&
	// (len(act.Nickname) < 1) &&
	// (len(act.Name) < 1)
}

// 如果所有这些关键数据都没有，则认为是 empty account
func (act *Account) IsValid() error {
	myValidator := validator.New()
	return myValidator.Struct(act)
}

func (act *Account) HasMboile() bool {
	return (len(act.Mobile) == constants.LengthMobileNumber)
}

func (act *Account) AssignFields(assign *Account) *Account {
	if (assign == nil) || (act.Aid != assign.Aid) {
		return nil
	}

	changed := false
	// 每一个string的有效性都必须做判断，因为 assign.IsValid() 中，一个 name == “” 也认为是有效的
	if (len(assign.Nickname) >= constants.MinLength_accountName) && (act.Nickname != assign.Nickname) {
		act.Nickname = assign.Nickname
		changed = true
	}

	if (len(assign.Mobile) >= constants.LengthMobileNumber) && (act.Mobile != assign.Mobile) {
		act.Mobile = assign.Mobile
		changed = true
	}

	if (len(assign.AvatarUrl) >= constants.MinLength_avatar_url) && (act.AvatarUrl != assign.AvatarUrl) {
		act.AvatarUrl = assign.AvatarUrl
		changed = true
	}

	if changed {
		return act
	}

	return nil
}
