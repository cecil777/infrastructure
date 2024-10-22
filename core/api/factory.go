package api

import "reflect"

var (
	endpointAPITypes = make(map[string]map[string]reflect.Type)
	invalid          = new(invalidAPI)
)

// Build is 创建IAPI
func Build(endpoint, name string) IAPI {
	if apiTypes, ok := endpointAPITypes[endpoint]; ok {
		if apiType, ok := apiTypes[name]; ok {
			instance := reflect.New(apiType).Interface().(IAPI)
			return instance
		}
	}
	return invalid
}

// Register is 注册api
func Register(endpoint, name string, api IAPI) {
	if _, ok := endpointAPITypes[endpoint]; !ok {
		endpointAPITypes[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	endpointAPITypes[endpoint][name] = apiType
}

// todo 完善 go reflect 反射文档
// todo 完善 go mock 文档 (mock 文件生成语句 mockgen -destination="i-api-mock.go" -package="api" -source .\i-api.go IAPI)
