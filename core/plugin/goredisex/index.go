package goredisex

import (
	"core/plugin/redisex"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

type redisi struct {
	host     string
	password string
	client   *redis.Client
}

func (r *redisi) Close() error {
	client := r.getClient()
	err := client.Close()

	return err
}

func (r *redisi) Del(arg ...string) (int, error) {
	client := r.getClient()
	del := client.Del(arg...)

	return int(del.Val()), del.Err()
}

func (r *redisi) Exists(str string) (bool, error) {
	client := r.getClient()
	res := client.Exists(str)

	if res.Val() == 1 {
		return true, res.Err()
	}

	return false, res.Err()
}

func (r *redisi) Get(str string) (string, error) {
	client := r.getClient()
	val := client.Get(str)

	return val.Result()
}

func (r *redisi) Set(key, val string, arg ...interface{}) (bool, error) {
	client := r.getClient()

	if len(arg) == 1 {
		if arg[0] == "NX" {
			res := client.SetNX(key, val, 0)
			return res.Result()
		} else if arg[0] == "XX" {
			res := client.SetXX(key, val, 0)
			return res.Result()
		} else {
			return false, errors.New("参数错误")
		}
	} else if len(arg) == 2 {
		if arg[0] == "EX" {
			s, ok := arg[1].(int)
			if ok {
				res := client.Set(key, val, time.Duration(s*1000000000))
				if res.Val() != "" {
					return true, res.Err()
				}
				return false, res.Err()
			}
		} else if arg[0] == "PX" {
			s, ok := arg[1].(int)
			if ok {
				res := client.Set(key, val, time.Duration(s*1000000))
				if res.Val() != "" {
					return true, res.Err()
				}
				return false, res.Err()
			}
		} else {
			return false, errors.New("参数错误")
		}
	} else if len(arg) == 3 {
		if arg[2] == "NX" {
			s, ok := arg[1].(int)
			if ok {
				if arg[0] == "EX" {
					res := client.SetNX(key, val, time.Duration(s*1000000000))
					return res.Result()
				} else if arg[0] == "PX" {
					res := client.SetNX(key, val, time.Duration(s*1000000))
					return res.Result()
				} else {
					return false, errors.New("参数错误")
				}
			}
		} else if arg[2] == "XX" {
			s, ok := arg[1].(int)
			if ok {
				if arg[0] == "EX" {
					res := client.SetXX(key, val, time.Duration(s*1000000000))
					return res.Result()
				} else if arg[0] == "PX" {
					res := client.SetXX(key, val, time.Duration(s*1000000))
					return res.Result()
				} else {
					return false, errors.New("参数错误")
				}
			}
		} else {
			return false, errors.New("参数错误")
		}
	} else if len(arg) == 0 {
		res := client.Set(key, val, 0)
		if res.Val() != "" {
			return true, res.Err()
		}
		return false, res.Err()
	}
	return false, errors.New("参数错误")
}

func (r *redisi) Time() (time.Time, error) {
	client := r.getClient()
	res := client.Time()

	return res.Result()
}

func (r *redisi) TTL(key string) (time.Duration, error) {
	client := r.getClient()
	ttl := client.TTL(key)

	return ttl.Result()
}

func (r *redisi) getClient() redis.Client {
	if r.client == nil {
		r.client = redis.NewClient(&redis.Options{
			Addr:     r.host,
			Password: r.password,
		})
	}

	return *r.client
}

func NewRedis(host, password string) redisex.IRedis {
	r := redisi{
		host:     host,
		password: password,
	}
	return &r
}
