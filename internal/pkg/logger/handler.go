package logger

type Handler interface {
	IsHandling(record Record) bool
	Handle(record Record) error
}
