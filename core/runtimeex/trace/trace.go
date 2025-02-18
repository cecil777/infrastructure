package trace

import (
	"github.com/cecil777/infrastructure/core/object"
	"github.com/cecil777/infrastructure/core/timeex"
)

type trace struct {
	id              string
	nowTime         timeex.INowTime
	onSpanEndAction func(labels map[string]interface{})
	stringGenerator object.IStringGenerator
}

func (m trace) BeginSpan(name, parentID string) ISpan {
	span := &span{
		labels:          make(map[string]interface{}),
		nowTime:         m.nowTime,
		onEndAction:     m.onSpanEndAction,
		stringGenerator: m.stringGenerator,
	}
	span.AddLabel("name", name)
	if parentID != "" {
		span.AddLabel("parentID", parentID)
	}
	span.AddLabel("traceID", m.GetID())
	nanoUnix := m.nowTime.UnixNano()
	span.AddLabel("beganOn", nanoUnix/1000/1000)
	return span
}

func (m *trace) GetID() string {
	if m.id == "" {
		m.id = m.stringGenerator.Generate()
	}

	return m.id
}
