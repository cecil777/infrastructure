package goredisex

import (
	"core/plugin/redisex"
	"time"

	"github.com/go-redis/redis"
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

func (r *redisi) Set(key, value string, extraArgs ...interface{}) (ok bool, err error) {
	var res string
	if len(extraArgs) == 0 {
		res, err = r.client.Set(key, value, 0).Result()
		ok = res == "OK"
	} else if len(extraArgs) == 1 {
		if extraArgs[0] == "nx" {
			ok, err = r.client.SetNX(key, value, 0).Result()
		} else {
			ok, err = r.client.SetXX(key, value, 0).Result()
		}
	} else if len(extraArgs) == 2 {
		res, err = r.client.Set(
			key,
			value,
			extraArgs[1].(time.Duration),
		).Result()
		ok = res == "OK"
	} else {
		if extraArgs[2] == "nx" {
			ok, err = r.client.SetNX(
				key,
				value,
				extraArgs[1].(time.Duration),
			).Result()
		} else {
			ok, err = r.client.SetXX(
				key,
				value,
				extraArgs[1].(time.Duration),
			).Result()
		}
	}
	return
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
