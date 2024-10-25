package log

type ILog interface {
	AddLabel(key, format string, v ...interface{}) ILog
	Debug()
	Error()
	Fatal()
	Info()
	Warning()
}
