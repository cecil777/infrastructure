package api

// Response is api响应结构
type Response struct {
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
}
