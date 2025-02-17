package logrusex

import (
	"encoding/json"
	"fmt"

	"github.com/cecil777/infrastructure/core/log"

	"github.com/sirupsen/logrus"
)

type logAdapter struct {
	labels map[string]string
}

type logFormat struct {
}

func (l *logAdapter) AddLabel(key, format string, v ...interface{}) log.ILog {
	l.labels[key] = fmt.Sprintf(format, v...)
	return l
}

func (l *logAdapter) Debug() {
	logrus.Debugln(mapToString(l.labels))
}

func (l *logAdapter) Error() {
	logrus.Error(mapToString(l.labels))
}

func (l *logAdapter) Fatal() {
	logrus.Fatal(mapToString(l.labels))
}

func (l *logAdapter) Info() {
	logrus.Info(mapToString(l.labels))
}

func (l *logAdapter) Warning() {
	logrus.Warning(mapToString(l.labels))
}

func mapToString(v interface{}) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

func (s logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] - %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
	return []byte(msg), nil
}

func NewLog() *logAdapter {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(new(logFormat))
	return &logAdapter{
		labels: map[string]string{},
	}
}
