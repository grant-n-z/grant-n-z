package log

import (
	"log"
	"os"
)

var Logger = NewLogger()

type Log struct {
	i *log.Logger
	w *log.Logger
	e *log.Logger
}

func NewLogger() Log {
	return Log{
		i: log.New(os.Stdout, "[INFO]", log.LstdFlags|log.LUTC),
		w: log.New(os.Stderr, "[WARN]", log.LstdFlags|log.LUTC),
		e: log.New(os.Stderr, "[ERROR]", log.LstdFlags|log.LUTC),
	}
}

func (l Log) Info(message ...interface{}) {
	l.i.Println(message...)
}

func (l Log) Warn(message ...interface{}) {
	l.w.Println(message...)
}

func (l Log) Error(message ...interface{}) {
	l.e.Println(message...)
}
