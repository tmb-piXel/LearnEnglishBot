package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
}

func Println(args ...interface{}) {
	Log.Print(args...)
}

func Printf(format string, args ...interface{}) {
	arguments := args[2:]
	Log.WithFields(logrus.Fields{
		"chatID":   args[0],
		"fullName": args[1],
	}).Printf(format, arguments...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	Log.Panic(args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}
