package randex

import (
	"math/rand"
	"time"
)

type stringGenerator struct {
	len        int
	sourceByte []byte
}

func (s *stringGenerator) Generate() string {
	bytes := make([]byte, s.len, s.len)
	allLen := len(s.sourceByte)

	rand.Seed(time.Now().UnixNano()) //确保种子随机性
	for i := 0; i < s.len; i++ {
		num := rand.Intn(allLen)
		bytes[i] = s.sourceByte[num]
	}

	return string(bytes)
}

func NewStringGenerator(b []byte, len int) *stringGenerator {
	return &stringGenerator{
		len:        len,
		sourceByte: b,
	}
}
