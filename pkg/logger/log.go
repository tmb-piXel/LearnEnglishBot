package logger

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

func init() {
	l = logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
}

func Println(args ...interface{}) {
	l.Print(args...)
}

func Printf(format string, args ...interface{}) {
	arguments := args[2:]
	l.WithFields(logrus.Fields{
		"chatID":   args[0],
		"fullName": args[1],
	}).Printf(format, arguments...)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	arguments := args[2:]
	l.WithFields(logrus.Fields{
		"chatID":   args[0],
		"fullName": args[1],
	}).Errorf(format, arguments...)
}

func Panic(args ...interface{}) {
	l.Panic(args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args...)
}
