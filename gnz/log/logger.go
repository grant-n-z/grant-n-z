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
		d:     NewDebugLog(),
		i:     NewInfoLog(),
		w:     NewWarnLog(),
		e:     NewErrorLog(),
	}
}

func NewDebugLog() *log.Logger {
	return log.New(os.Stdout, "[D]", log.LstdFlags|log.LUTC)
}

func NewInfoLog() *log.Logger {
	return log.New(os.Stdout, "[I]", log.LstdFlags|log.LUTC)
}

func NewWarnLog() *log.Logger {
	return log.New(os.Stderr, "[W]", log.LstdFlags|log.LUTC)
}

func NewErrorLog() *log.Logger {
	return log.New(os.Stderr, "[E]", log.LstdFlags|log.LUTC)
}

func (l Log) Debug(log ...interface{}) {
	if strings.EqualFold(l.level, "DEBUG") || strings.EqualFold(l.level, "debug") {
		_, file, line, _ := runtime.Caller(1)
		execFile := strings.Split(file, "/")
		data := fmt.Sprintf("%s/%s:%v", execFile[len(execFile)-2], execFile[len(execFile)-1], line)
		l.d.Println(data, fmt.Sprint(log...))
	}
}

func (l Log) Info(log ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	execFile := strings.Split(file, "/")
	data := fmt.Sprintf("%s/%s:%v", execFile[len(execFile)-2], execFile[len(execFile)-1], line)
	l.i.Println(data, fmt.Sprint(log...))
}

func (l Log) Warn(log ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	execFile := strings.Split(file, "/")
	data := fmt.Sprintf("%s/%s:%v", execFile[len(execFile)-2], execFile[len(execFile)-1], line)
	l.w.Println(data, fmt.Sprint(log...))
}

func (l Log) Error(log ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	execFile := strings.Split(file, "/")
	data := fmt.Sprintf("%s/%s:%v", execFile[len(execFile)-2], execFile[len(execFile)-1], line)
	l.e.Println(data, fmt.Sprint(log...))
}

func (l Log) Trace(log interface{}) {
	l.e.Println(log)
	for depth := 0; ; depth++ {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		fmt.Printf("%d: %v:%d\n", depth, file, line)
	}
}
