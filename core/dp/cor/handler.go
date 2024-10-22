package cor

type handler struct {
	isBreak bool
	nexts   []interface{}
}

func (m *handler) Break() {
	m.isBreak = true
}

func (m handler) Handle() (err error) {
	if m.isBreak {
		return
	}

	for _, r := range m.nexts {
		h := r.(IHandler)
		if err = h.Handle(); err != nil || h.IsBreak() {
			break
		}
	}

	return
}

func (m handler) IsBreak() bool {
	return m.isBreak
}

func (m *handler) SetNext(handler IHandler) IHandler {
	m.nexts = append(m.nexts, handler)
	return m
}

// New is 创建IHandler
func New() IHandler {
	return &handler{
		nexts: make([]interface{}, 0),
	}
}
