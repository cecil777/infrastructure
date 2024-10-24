package goredisex

import (
	"time"

	"github.com/cecil777/infrastructure/core/plugin/redisex"
	"github.com/go-redis/redis"
)

type redisAdapter struct {
	host     string
	password string
	client   *redis.Client
}

func (r *redisAdapter) Close() error {

	return r.getClient().Close()
}

func (r *redisAdapter) Del(arg ...string) (int, error) {
	del := r.getClient().Del(arg...)

	return int(del.Val()), del.Err()
}

func (r *redisAdapter) Exists(str string) (bool, error) {
	res := r.getClient().Exists(str)

	return res.Val() == 1, res.Err()
}

func (r *redisAdapter) Get(str string) (string, error) {

	return r.getClient().Get(str).Result()
}

func (r *redisAdapter) Set(key, value string, extraArgs ...interface{}) (ok bool, err error) {
	var res string
	if len(extraArgs) == 0 {
		res, err = r.getClient().Set(key, value, 0).Result()
		ok = res == "OK"
	} else if len(extraArgs) == 1 {
		if extraArgs[0] == "nx" {
			ok, err = r.getClient().SetNX(key, value, 0).Result()
		} else if extraArgs[0] == "xx" {
			ok, err = r.getClient().SetXX(key, value, 0).Result()
		} else {
			panic("redis set 参数有误")
		}
	} else if len(extraArgs) == 2 {
		s, okk := extraArgs[1].(int)
		if okk {
			if extraArgs[0] == "ex" {
				s = s * 1000 * 1000 * 1000
			} else if extraArgs[0] == "px" {
				s = s * 1000 * 1000
			} else {
				panic("redis set 参数有误")
			}
			res, err = r.getClient().Set(
				key,
				value,
				time.Duration(s),
			).Result()
			ok = res == "OK"
		} else {
			panic("redis set 参数有误")
		}
	} else if len(extraArgs) == 3 {
		s, okk := extraArgs[1].(int)
		if okk {
			if extraArgs[0] == "ex" {
				s = s * 1000 * 1000 * 1000
			} else if extraArgs[1] == "px" {
				s = s * 1000 * 1000
			} else {
				panic("redis set 参数有误")
			}
			if extraArgs[2] == "nx" {
				ok, err = r.getClient().SetNX(
					key,
					value,
					time.Duration(s),
				).Result()
			} else if extraArgs[2] == "xx" {
				ok, err = r.getClient().SetXX(
					key,
					value,
					time.Duration(s),
				).Result()
			} else {
				panic("redis set 参数有误")
			}
		} else {
			panic("redis set 参数错误")
		}
	} else {
		panic("redis set 参数过多")
	}
	return
}

func (r *redisAdapter) Time() (time.Time, error) {

	return r.getClient().Time().Result()
}

func (r *redisAdapter) TTL(key string) (time.Duration, error) {

	return r.getClient().TTL(key).Result()
}

func (r *redisAdapter) getClient() *redis.Client {
	if r.client == nil {
		r.client = redis.NewClient(&redis.Options{
			Addr:     r.host,
			Password: r.password,
		})
	}

	return r.client
}

func NewRedis(host, password string) redisex.IRedis {
	return &redisAdapter{
		host:     host,
		password: password,
	}
}
