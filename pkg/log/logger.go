package log

type Logger interface {
	Emergency(message string, fields ...Field)
	Alert(message string, fields ...Field)
	Critical(message string, fields ...Field)
	Error(message string, fields ...Field)
	Warning(message string, fields ...Field)
	Notice(message string, fields ...Field)
	Info(message string, fields ...Field)
	Debug(message string, fields ...Field)
}
