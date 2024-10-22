package goredisex

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	self   = NewRedis("127.0.0.1:6379", "")
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
)

func Test_goRedis_Del(t *testing.T) {
	key := "Test_goRedis_Del"
	val := "test"
	client.Set(key, val, 0).Result()

	defer client.Del(key)

	get, err2 := self.Get(key)
	assert.Equal(t, get, val)
	assert.Equal(t, err2, nil)
	count, err := self.Del(key)
	assert.Equal(t, count, 1)
	assert.Equal(t, err, nil)
	get, err2 = self.Get(key)
	assert.Equal(t, get, "")
}

func Test_goRedis_SetXXNX(t *testing.T) {
	key1 := "a"
	val1 := "test1"
	set, err := self.Set(key1, val1, "XX")
	assert.Equal(t, set, false)
	assert.Equal(t, err, nil)
	get, err := self.Get(key1)
	assert.Equal(t, get, "")

	set, err = self.Set(key1, val1, "NX")
	assert.Equal(t, set, true)
	assert.Equal(t, err, nil)
	get, err = self.Get(key1)
	assert.Equal(t, get, val1)
	assert.Equal(t, err, nil)

	set, err = self.Set(key1, val1, "XX")
	assert.Equal(t, set, true)
	assert.Equal(t, err, nil)
	get, err = self.Get(key1)
	assert.Equal(t, get, val1)
}

func Test_goRedis_SetEXPX(t *testing.T) {
	key1 := "b"
	val1 := "test2"
	second := 60
	set, err := self.Set(key1, val1, "EX", second)
	assert.Equal(t, set, true)
	assert.Equal(t, err, nil)
	ttl, err := self.TTL(key1)
	assert.Equal(t, ttl, time.Duration(second*1000*1000*1000))

	key2 := "c"
	val2 := "test3"
	millisecond := 10000
	set, err = self.Set(key2, val2, "PX", millisecond)
	assert.Equal(t, set, true)
	assert.Equal(t, err, nil)
	ttl, err = self.TTL(key2)
	assert.Equal(t, ttl, time.Duration(millisecond*1000*1000))
}

func Test_goRedis_Exists(t *testing.T) {
	key := "d"
	val := "test4"

	set, err := self.Set(key, val)
	assert.Equal(t, set, true)
	assert.Equal(t, err, nil)
	exists, err := self.Exists(key)
	assert.Equal(t, exists, true)
	assert.Equal(t, err, nil)
}

func Test_goRedis_Time(t *testing.T) {
	t2, err := self.Time()
	assert.Equal(t, err, nil)
	fmt.Println(t2)
}
