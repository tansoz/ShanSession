package Session

import (
	"crypto/md5"
	"fmt"
	"github.com/tansoz/ShanMap"
	"math/rand"
	"net/http"
	"time"
)

type SessionManagerImpl struct {
	SessionManager
	list          ShanMap.ShanMap
	sessionIdName string
	expire        int64
}

func NewSessionManager(sessionIdName string, expire int64) SessionManager {

	SM := &SessionManagerImpl{
		list:          ShanMap.NewSafeMap(),
		sessionIdName: sessionIdName,
		expire:        expire,
	}
	if expire >= 0 {
		SM.ScanSession()
	}

	return SM

}

// Http 获取当前用户的session信息
func (this *SessionManagerImpl) HttpSession(rw http.ResponseWriter, r *http.Request) Session {

	sessId := this.MakeSessionId()
	needCookie := true // 是否需要set-cookie

	if c, err := r.Cookie(this.sessionIdName); err == nil {

		sessId = c.Value
		needCookie = false
	}

	sess := this.GetSession(sessId)
	if needCookie {
		this.Cookie(sess, rw)
	}

	return sess

}

// 根据session id 获取session 信息
func (this *SessionManagerImpl) GetSession(sessionId string) Session {
	if v := this.list.Get(sessionId).Interface(); v == nil {

		return this.MakeSession(sessionId)

	} else {

		if v.(Session).IsExpired() {
			return this.MakeSession(sessionId)
		} else {
			v.(Session).UpdateExpired()
			return v.(Session)
		}
	}

}

func (this *SessionManagerImpl) MakeSession(sessionId string) Session {

	session := NewSession(sessionId, this.expire)

	this.list.Set(session.GetSessionId(), session)

	return session

}

func (this *SessionManagerImpl) Cookie(session Session, rw http.ResponseWriter) {

	http.SetCookie(rw, &http.Cookie{
		Name:     this.sessionIdName,
		Value:    session.GetSessionId(),
		Path:     "/",
		HttpOnly: true,
	})

}
func (this *SessionManagerImpl) GetSessionIdName() string {
	return this.sessionIdName
}
func (this *SessionManagerImpl) MakeSessionId() string {
	return md5String(fmt.Sprintf("%s#%s@%f", this.sessionIdName, time.Now().String(), rand.Float32()))
}
func md5String(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func (this *SessionManagerImpl) ScanSession() {

	go func() {

		for {

			for k, i := range this.list.Map() {
				if i.Interface().(Session).IsExpired() {
					this.list.Del(k)
				}

				time.Sleep(20 * time.Millisecond)
			}
			time.Sleep(20 * time.Second)
		}

	}()

}
