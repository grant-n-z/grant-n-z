package log

import (
	"log"
	"os"
	"strings"
)

var Logger Log

type Log struct {
	level string
	d     *log.Logger
	i     *log.Logger
	w     *log.Logger
	e     *log.Logger
}

func InitLogger(logLevel string) {
	Logger = Log{
		level: logLevel,
		d:     NewLoglevelDebug(),
		i:     NewLoglevelInfo(),
		w:     NewLoglevelWarn(),
		e:     NewLoglevelError(),
	}
}

func NewLoglevelDebug() *log.Logger {
	return log.New(os.Stdout, "[DEBUG]", log.LstdFlags|log.LUTC)
}

func NewLoglevelInfo() *log.Logger {
	return log.New(os.Stdout, "[INFO]", log.LstdFlags|log.LUTC)
}

func NewLoglevelWarn() *log.Logger {
	return log.New(os.Stderr, "[WARN]", log.LstdFlags|log.LUTC)
}

func NewLoglevelError() *log.Logger {
	return log.New(os.Stderr, "[ERROR]", log.LstdFlags|log.LUTC)
}

func (l Log) Fatal(v ...interface{}) {
	log.Fatal(v)
}

func (l Log) Debug(log ...interface{})  {
	if strings.EqualFold(l.level, "DEBUG") || strings.EqualFold(l.level, "debug") {
		l.d.Println(log...)
	}
}

func (l Log) Info(log ...interface{}) {
	l.i.Println(log...)
}

func (l Log) Warn(log ...interface{}) {
	l.w.Println(log...)
}

func (l Log) Error(log ...interface{}) {
	l.e.Println(log...)
}
