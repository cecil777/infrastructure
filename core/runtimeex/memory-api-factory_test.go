package runtimeex

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_memoryAPIFactory_Build(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		endpoint := "endpoint"
		name := "name"
		self := make(memoryAPIFactory)
		res := self.Build(endpoint, name)
		assert.Equal(t, res, invalidAPISingleton)
	})
}

func Test_memoryAPIFactory_Register(t *testing.T) {
	endpoint := "endpoint"
	name := "name"
	self := make(memoryAPIFactory)
	self.Register(endpoint, name, invalidAPISingleton)

	apiTypes, ok := self[endpoint]
	assert.True(t, ok)

	apiType, ok := apiTypes[name]
	assert.True(t, ok)
	assert.Equal(
		t,
		apiType,
		reflect.TypeOf(invalidAPI{}),
	)
}
