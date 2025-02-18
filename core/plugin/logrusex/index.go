package logrusex

import (
	"fmt"

	"github.com/cecil777/infrastructure/core/log"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type logAdapter struct {
	assertAction func(string)
	labels       map[string]string
}

func (l *logAdapter) AddLabel(key, format string, v ...interface{}) log.ILog {
	l.labels[key] = fmt.Sprintf(format, v...)
	return l
}

func (l *logAdapter) Debug() {
	l.log(logrus.Debug)
}

func (l *logAdapter) Error() {
	l.log(logrus.Error)
}

func (l *logAdapter) Fatal() {
	l.log(logrus.Fatal)
}

func (l *logAdapter) Info() {
	l.log(logrus.Info)
}

func (l *logAdapter) Warning() {
	l.log(logrus.Warning)
}

func (l *logAdapter) log(action func(args ...interface{})) {
	marshal, err := jsoniter.MarshalToString(l.labels)
	if err != nil {
		fmt.Println(err)
	}

	if l.assertAction != nil {
		l.assertAction(marshal)
	} else {
		action(marshal)
	}
}

type logFormat struct {
}

func (s logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] - %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
	return []byte(msg), nil
}

func NewLog() log.ILog {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(new(logFormat))
	return &logAdapter{
		labels: map[string]string{},
	}
}
