package ioex

// IPath is 路径接口
type IPath interface {
	Join(paths ...string) string
}
