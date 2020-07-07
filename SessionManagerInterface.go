package Session

import "net/http"

type SessionManager interface {
	GetSession(sessionId string) Session
	HttpSession(rw http.ResponseWriter, r *http.Request) Session
	MakeSession(sessionId string) Session
	Cookie(session Session, rw http.ResponseWriter)
	MakeSessionId() string
	GetSessionIdName() string
}
