package logrusex

import (
	"testing"
)

func TestInfo(t *testing.T) {
	log := NewLog()
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test111", 1, 3, true)
	log.Info()
}

func TestDebug(t *testing.T) {
	log := NewLog()
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Debug()
}

func TestError(t *testing.T) {
	log := NewLog()
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Error()
}

func TestWarning(t *testing.T) {
	log := NewLog()
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Warning()
}

func TestFatal(t *testing.T) {
	log := NewLog()
	log.AddLabel("a", "test=%v, num1=%d, num2=%d, bool=%v", "test1", 1, 3, true)
	log.Fatal()
}
