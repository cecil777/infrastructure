package trace

import (
	"github.com/cecil777/infrastructure/core/object"
	"github.com/cecil777/infrastructure/core/timeex"
)

type factory struct {
	nowTime         timeex.INowTime
	onSpanEndAction func(labels map[string]interface{})
	stringGenerator object.IStringGenerator
}

func (m factory) Build(traceID string) ITrace {
	return &trace{
		id:              traceID,
		nowTime:         m.nowTime,
		onSpanEndAction: m.onSpanEndAction,
		stringGenerator: m.stringGenerator,
	}
}

// NewFactory is 跟踪工厂
func NewFactory(nowTime timeex.INowTime, onSpanEndAction func(labels map[string]interface{}), stringGenerator object.IStringGenerator) IFactory {
	return &factory{
		nowTime:         nowTime,
		onSpanEndAction: onSpanEndAction,
		stringGenerator: stringGenerator,
	}
}
