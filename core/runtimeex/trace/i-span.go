package trace

// ISpan is 段接口
type ISpan interface {
	AddLabel(key string, value interface{})
	End()
	GetID() string
}
