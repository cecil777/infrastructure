package runtimeex

import (
	"reflect"

	"github.com/cecil777/infrastructure/core/api"
	"github.com/cecil777/infrastructure/core/errorex"
)

var invalidAPISingleton = new(invalidAPI)

type invalidAPI struct{}

func (m invalidAPI) Call() (interface{}, error) {
	return nil, errorex.New(errorex.APICode, "")
}

type memoryAPIFactory map[string]map[string]reflect.Type

func (m memoryAPIFactory) Build(endpoint, name string) api.IAPI {
	if apiTypes, ok := m[endpoint]; ok {
		if apiType, ok := apiTypes[name]; ok {
			instance := reflect.New(apiType).Interface().(api.IAPI)
			return instance
		}
	}

	return invalidAPISingleton
}

func (m memoryAPIFactory) Register(endpoint, name string, api api.IAPI) {
	if _, ok := m[endpoint]; !ok {
		m[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	m[endpoint][name] = apiType
}

// NewAPIFactory is 内存api工厂
func NewAPIFactory() api.IFactory {
	return make(memoryAPIFactory)
}
