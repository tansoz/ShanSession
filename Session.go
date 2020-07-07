package Session

import (
	"github.com/tansoz/ShanMap"
	"time"
)

type SessionImpl struct {
	Session
	data      ShanMap.ShanMap
	expired   int64
	expire    int64
	sessionId string
}

func NewSession(sessionId string, expire int64) Session {

	time := time.Now().Unix() + expire
	if expire < 0 {
		time = -1
	}

	return &SessionImpl{
		sessionId: sessionId,
		expired:   time,
		expire:    expire,
		data:      ShanMap.NewSafeMap(),
	}
}

func (this *SessionImpl) GetSessionId() string {
	return this.sessionId
}
func (this *SessionImpl) Data() ShanMap.ShanMap {
	return this.data
}
func (this *SessionImpl) IsExpired() bool {
	if this.expired > time.Now().Unix() || this.expire < 0 {
		return false
	} else {
		return true
	}
}
func (this *SessionImpl) UpdateExpired() {
	if this.expire != -1 {
		if this.expired > time.Now().Unix() {
			this.expired += this.expire - (this.expired - time.Now().Unix())
		}
	}

}
func (this *SessionImpl) Refresh() {
	this.UpdateExpired()
	this.data = ShanMap.NewSafeMap()
}
