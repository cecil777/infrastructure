package goredisex

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
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

	res, err := self.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, res, val)

	count, err := self.Del(key)
	assert.NoError(t, err)
	assert.Equal(t, count, 1)

	res, err = self.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, res, "")
}

func Test_goRedis_SetXXNX(t *testing.T) {
	key1 := "a"
	val1 := "test1"
	set, err := self.Set(key1, val1, "xx")
	defer self.Del(key1)
	assert.NoError(t, err)
	assert.Equal(t, set, false)

	res, err := self.Get(key1)
	assert.NoError(t, err)
	assert.Equal(t, res, "")

	set, err = self.Set(key1, val1, "nx")
	assert.NoError(t, err)
	assert.Equal(t, set, true)

	res, err = self.Get(key1)
	assert.NoError(t, err)
	assert.Equal(t, res, val1)

	set, err = self.Set(key1, val1, "xx")
	assert.NoError(t, err)
	assert.Equal(t, set, true)

	res, err = self.Get(key1)
	assert.NoError(t, err)
	assert.Equal(t, res, val1)

}

func Test_goRedis_SetEXPX(t *testing.T) {
	key1 := "b"
	val1 := "test2"
	second := 60
	set, err := self.Set(key1, val1, "ex", second)
	assert.NoError(t, err)
	assert.Equal(t, set, true)

	ttl, err := self.TTL(key1)
	assert.NoError(t, err)
	assert.Equal(t, ttl, time.Duration(second*1000*1000*1000))

	key2 := "c"
	val2 := "test3"
	millisecond := 10000
	set, err = self.Set(key2, val2, "px", millisecond)
	assert.NoError(t, err)
	assert.Equal(t, set, true)

	ttl, err = self.TTL(key2)
	assert.NoError(t, err)
	assert.Equal(t, ttl, time.Duration(millisecond*1000*1000))
}

func Test_goRedis_Exists(t *testing.T) {
	key := "d"
	val := "test4"

	set, err := self.Set(key, val)
	defer self.Del(key)
	assert.NoError(t, err)
	assert.Equal(t, set, true)

	exists, err := self.Exists(key)
	assert.NoError(t, err)
	assert.Equal(t, exists, true)

}

func Test_goRedis_Time(t *testing.T) {
	t2, err := self.Time()
	assert.NoError(t, err)
	fmt.Println(t2)
}
