package logger

import "fmt"

var Verbose = true

func Debugf(format string, args ...interface{}) {
	if !Verbose {
		return
	}

	printf(format, args...)
}

func Debugln(s string) {
	Debugf(s)
}

func Infof(format string, args ...interface{}) {
	if !Verbose {
		return
	}

	printf(format, args...)
}

func Infoln(s string) {
	Infof(s)
}

func Warnf(format string, args ...interface{}) {
	printf(format, args...)
}

func Warnln(s string) {
	Warnf(s)
}

func Errorf(format string, args ...interface{}) {
	printf(format, args...)
}

func Errorln(s string) {
	Errorf(s)
}

func printf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
