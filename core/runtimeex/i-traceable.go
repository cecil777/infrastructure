package runtimeex

// ITraceable is 追踪借口
type ITraceable interface {
	SetTraceID(string)
	SetTraceSpanID(string)
}
