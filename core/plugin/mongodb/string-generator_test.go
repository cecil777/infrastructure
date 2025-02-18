package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectID2String(t *testing.T) {
	t.Run("生成2个", func(t *testing.T) {
		generator := NewStringGenerator()
		s1 := generator.Generate()
		s2 := generator.Generate()
		assert.NotEqual(t, s1, s2)
	})
}
