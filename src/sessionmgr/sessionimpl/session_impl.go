package sessionimpl

import (
	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
)

const (
	// length of sessio token, that is a UUID without '-' symbol
	Length_of_SessionToken = 32
)

var gSession *sessionManager

type sessionManager struct {
	session map[string]*sessiondata.SessionInfo
}

func init() {
	gSession = &sessionManager{
		session: make(map[string]*sessiondata.SessionInfo),
	}
}

func AddSession(sessionID string, sinfo *sessiondata.SessionInfo) {
	gSession.session[sessionID] = sinfo
}

func GetSession(sessionID string) *sessiondata.SessionInfo {
	if Length_of_SessionToken != len(sessionID) {
		return nil
	}
	return gSession.session[sessionID]
}

func RemoveSession(sessionID string) {
	delete(gSession.session, sessionID)
}

func IsSessionExist(sessionID string) bool {
	_, ok := gSession.session[sessionID]
	return ok
}
