package sessionmgr

import (
	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
	"lizisky.com/lizisky/src/sessionmgr/sessionimpl"
)

// make sessionmgr as simple as possiblez
// it simliar as interface

func AddSession(sessionID string, sinfo *sessiondata.SessionInfo) {
	sessionimpl.AddSession(sessionID, sinfo)
}

func GetSession(sessionID string) *sessiondata.SessionInfo {
	return sessionimpl.GetSession(sessionID)
}

func RemoveSession(sessionID string) {
	sessionimpl.RemoveSession(sessionID)
}

func IsSessionExist(sessionID string) bool {
	return sessionimpl.IsSessionExist(sessionID)
}
