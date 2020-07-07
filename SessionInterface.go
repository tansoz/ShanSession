package Session

import "github.com/tansoz/ShanMap"

type Session interface {
	GetSessionId() string
	Data() ShanMap.ShanMap
	IsExpired() bool // 是否到期
	UpdateExpired()
	Refresh()
}
