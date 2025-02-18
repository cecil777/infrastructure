package trace

type ITrace interface {
	BeginSpan(name, parentID string) ISpan
	GetID() string
}
