package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
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
	return log.New(os.Stdout, "[DEBUG]", log.LstdFlags|log.Lshortfile|log.LUTC)
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
		_, file, line, _ := runtime.Caller(1)
		execFile := strings.Split(file, "/")
		logData := log[0]
		data := fmt.Sprintf("%s/%s:%v %s", execFile[len(execFile)-2], execFile[len(execFile)-1], line, logData)

		l.d.Println(data)
	}
}

func (l Log) Info(log ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	execFile := strings.Split(file, "/")
	logData := log[0]
	data := fmt.Sprintf("%s/%s:%v %s", execFile[len(execFile)-2], execFile[len(execFile)-1], line, logData)

	l.i.Println(data)
}

func (l Log) Warn(log ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	execFile := strings.Split(file, "/")
	logData := log[0]
	data := fmt.Sprintf("%s/%s:%v %s", execFile[len(execFile)-2], execFile[len(execFile)-1], line, logData)

	l.w.Println(data)
}

func (l Log) Error(log ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	execFile := strings.Split(file, "/")
	logData := log[0]
	data := fmt.Sprintf("%s/%s:%v %s", execFile[len(execFile)-2], execFile[len(execFile)-1], line, logData)

	l.e.Println(data)
}
