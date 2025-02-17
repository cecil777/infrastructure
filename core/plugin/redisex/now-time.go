package redisex

import "github.com/cecil777/infrastructure/core/timeex"

type nowTime struct {
	IRedis
}

func (m nowTime) Unix() int64 {
	t, err := m.Time()
	if err != nil {
		return 0
	}

	return t.Unix()
}

// UnixNano is 当前Unix纳秒级
func (m nowTime) UnixNano() int64 {
	t, err := m.Time()
	if err != nil {
		return 0
	}

	return t.UnixNano()
}

// NewNowTime is 当前时间实现类(redis)
func NewNowTime(redis IRedis) timeex.INowTime {
	return &nowTime{redis}
}
