package logger

import "fmt"

var log Logger

func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "console":
		log, err = NewConsole(config)
	case "file":
		log, err = NewFileLog(config)
	default:
		err = fmt.Errorf("unspport logger name%s", name)
	}
	return
}

func Debug(fmtstr string, args ...interface{}) {
	log.DeBug(fmtstr, args...)
}

func Trace(fmtstr string, args ...interface{}) {
	log.Trace(fmtstr, args...)
}

func Info(fmtstr string, args ...interface{}) {
	log.Info(fmtstr, args...)
}

func Warn(fmtstr string, args ...interface{}) {
	log.Warn(fmtstr, args...)
}

func Error(fmtstr string, args ...interface{}) {
	log.Error(fmtstr, args...)
}

func Fatal(fmtstr string, args ...interface{}) {
	log.Fatal(fmtstr, args...)
}
