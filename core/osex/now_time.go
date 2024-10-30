package osex

import (
	"time"

	"github.com/cecil777/infrastructure/core/timeex"
)

type nowTime struct{}

func (m nowTime) Unix() int64 {
	return time.Now().Unix()
}

func (m nowTime) UnixNano() int64 {
	return time.Now().UnixNano()
}

// NewNowTime is 创建系统timeex.INowTime
func NewNowTime() timeex.INowTime {
	return new(nowTime)
}
