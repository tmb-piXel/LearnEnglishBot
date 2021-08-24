package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	// if os.Getenv("environment") == "production" {
	Log.SetFormatter(&logrus.JSONFormatter{})
	f, _ := os.OpenFile("log", os.O_WRONLY|os.O_CREATE, 0755)
	Log.SetOutput(f)
	// }
}

func Println(args ...interface{}) {
	Log.Print(args...)
}

func Printf(format string, args ...interface{}) {
	Log.Printf(format, args...)
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
