package trace

import (
	"github.com/cecil777/infrastructure/core/object"
	"github.com/cecil777/infrastructure/core/timeex"
)

type span struct {
	id              string
	labels          map[string]interface{}
	nowTime         timeex.INowTime
	onEndAction     func(labels map[string]interface{})
	stringGenerator object.IStringGenerator
}

func (m *span) AddLabel(key string, value interface{}) {
	m.labels[key] = value
}

func (m span) End() {
	unixNano := m.nowTime.UnixNano()
	m.AddLabel("endedOn", unixNano/1000/1000)
	m.AddLabel("id", m.GetID())
	m.onEndAction(m.labels)
	m.labels = make(map[string]interface{})
}

func (m *span) GetID() string {
	if m.id == "" {
		m.id = m.stringGenerator.Generate()
	}

	return m.id
}
