package logger

type LogInterface interface {
	Error(err error)
	Warn(msg string)
	Info(msg string)
}
