package redisex

import "time"

// IRedis is redis接口
type IRedis interface {
	Close() error
	Del(...string) (int, error)
	Exists(string) (bool, error)
	Get(string) (string, error)
	Set(string, string, ...interface{}) (bool, error)
	Time() (time.Time, error)
	TTL(key string) (time.Duration, error)
}
