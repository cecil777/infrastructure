package trace

// IFactory is 跟踪工厂
type IFactory interface {
	Build(traceID string) ITrace
}
