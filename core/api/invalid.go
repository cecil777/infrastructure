package api

import "core/errorex"

// invalidAPI is 禁用(空)的API
type invalidAPI struct {
}

func (a invalidAPI) Call() (interface{}, error) {
	return nil, errorex.New(errorex.APICode, "")
}