package logrusex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	log := logAdapter{
		assertAction: func(text string) {
			assert.Equal(t, text, `{"a":"test=test1, num1=1, num2=3, bool=true"}`)
		},
		labels: make(map[string]string),
	}
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Info()
}

func TestDebug(t *testing.T) {
	log := logAdapter{
		assertAction: func(text string) {
			assert.Equal(t, text, `{"a":"test=test1, num1=1, num2=3, bool=true"}`)
		},
		labels: make(map[string]string),
	}
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Debug()
}

func TestError(t *testing.T) {
	log := logAdapter{
		assertAction: func(text string) {
			assert.Equal(t, text, `{"a":"test=test1, num1=1, num2=3, bool=true"}`)
		},
		labels: make(map[string]string),
	}
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Error()
}

func TestWarning(t *testing.T) {
	log := logAdapter{
		assertAction: func(text string) {
			assert.Equal(t, text, `{"a":"test=test1, num1=1, num2=3, bool=true"}`)
		},
		labels: make(map[string]string),
	}
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Warning()
}

func TestFatal(t *testing.T) {
	log := logAdapter{
		assertAction: func(text string) {
			assert.Equal(t, text, `{"a":"test=test1, num1=1, num2=3, bool=true"}`)
		},
		labels: make(map[string]string),
	}
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Fatal()
}
