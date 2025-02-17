package api

// IFactory is api工厂
type IFactory interface {
	Build(endpoint, name string) IAPI
	Register(endpoint, name string, api IAPI)
}
