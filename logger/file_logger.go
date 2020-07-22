package logger

import (
	"fmt"
	"os"
	"strconv"
)

type Filelog struct {
	level    int
	logPath  string
	logFile  string
	file     *os.File
	warnFile *os.File
}

func NewFileLog(msg map[string]string) Logger {
	var filelog *Filelog = &Filelog{}
	LogPath, ok := msg["log_path"]
	if !ok {
		LogPath = "/var/log/go.log"
	}
	filelog.logPath = LogPath
	Level, ok := msg["log_level"]
	if !ok {
		Level = "0"
	}
	level, err := strconv.Atoi(Level)
	if err != nil {
		level = 0
	}
	filelog.level = level
	LogFile, ok := msg["log_file"]
	if !ok {
		LogFile = "go.log"
	}
	filelog.logFile = LogFile
	filelog.init()
	return filelog
}

func (f *Filelog) init() {
	filename := fmt.Sprintf("%s/%s", f.logPath, f.logFile)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Sprintf("open logfile %s failed,Error:%v", filename, err))
	}
	f.file = file
	filename = fmt.Sprintf("%s/%s.wf", f.logPath, f.logFile)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Sprintf("open logfile %s failed,Error:%v", filename, err))
	}
	f.warnFile = file
}

func (f *Filelog) Setlevel(level int) {
	if level < LeDeBug || level > LeFatal {
		level = LeDeBug
	}
	f.level = level
}

func (f *Filelog) DeBug(fmtstr string, args ...interface{}) {
	if f.level > LeDeBug {
		return
	}
	writeLog(f.file, LeDeBug, fmtstr, args...)

}

func (f *Filelog) Trace(fmtstr string, args ...interface{}) {
	if f.level > LeTrace {
		return
	}
	writeLog(f.file, LeTrace, fmtstr, args...)
}

func (f *Filelog) Info(fmtstr string, args ...interface{}) {
	if f.level > LeInfo {
		return
	}
	writeLog(f.file, LeInfo, fmtstr, args...)
}

func (f *Filelog) Warn(fmtstr string, args ...interface{}) {
	if f.level > LeWarn {
		return
	}
	writeLog(f.file, LeWarn, fmtstr, args...)
}

func (f *Filelog) Error(fmtstr string, args ...interface{}) {
	if f.level > LeError {
		return
	}
	writeLog(f.warnFile, LeError, fmtstr, args...)
}

func (f *Filelog) Fatal(fmtstr string, args ...interface{}) {
	writeLog(f.warnFile, LeFatal, fmtstr, args...)
}

func (f *Filelog) Close() {
	f.file.Close()
	f.warnFile.Close()
}
