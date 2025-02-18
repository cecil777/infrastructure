package randex

import (
	"testing"

	"github.com/cecil777/go-underscore"
	"github.com/stretchr/testify/assert"
)

func TestStringGenerator_Generate(t *testing.T) {
	// str长度为6且字符只包含a-z
	sourceByte := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	length := 6
	stringGenerator := NewStringGenerator(sourceByte, length)

	//测试100次
	for i := 0; i < 100; i++ {
		str := stringGenerator.Generate()
		assert.Len(t, str, length) //长度断言
		//字符串每个字是否来自byte中

		assert.True(t, underscore.Chain([]byte(str)).All(func(r byte, _ int) bool {
			return underscore.Chain(sourceByte).Any(func(cr byte, _ int) bool {
				return cr == r
			})
		}))
	}

	// str为4位纯数字字符串
	sourceByte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	length = 4
	stringGenerator = NewStringGenerator(sourceByte, length)
	//测试100次
	for i := 0; i < 100; i++ {
		str := stringGenerator.Generate()
		assert.Len(t, str, length) //长度断言

		assert.True(t, underscore.Chain([]byte(str)).All(func(r byte, _ int) bool {
			return underscore.Chain(sourceByte).Any(func(cr byte, _ int) bool {
				return cr == r
			})
		}))
	}

}
