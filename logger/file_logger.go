package logger

import (
	"fmt"
	"os"
	"strconv"
)

type Filelog struct {
	level       int
	logPath     string
	logFile     string
	file        *os.File
	warnFile    *os.File
	logDataChan chan *LogData
}

func NewFileLog(msg map[string]string) (log Logger, err error) {
	var filelog *Filelog
	LogPath, ok := msg["log_path"]
	if !ok {
		LogPath = "/var/log/go.log"
	}
	Level, ok := msg["log_level"]
	if !ok {
		Level = "debug"
	}
	level := stringtoLevel(Level)
	LogFile, ok := msg["log_file"]
	if !ok {
		LogFile = "go.log"
	}
	logDataChan, ok := msg["log_chan_size"]
	if !ok {
		logDataChan = "50000"
	}
	chanNum, err := strconv.Atoi(logDataChan)
	if err != nil {
		chanNum = 50000
	}
	filelog.level = level
	filelog.logFile = LogFile
	filelog.logPath = LogPath
	filelog.logDataChan = make(chan *LogData, chanNum)
	log = filelog
	filelog.Init()
	return
}

func (f *Filelog) Init() {
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
	go f.writelog_d()
}

func (f *Filelog) writelog_d() {
	for logData := range f.logDataChan {
		var file = f.file
		if logData.WarnOrFatal {
			file = f.warnFile
		}
		fmt.Fprintf(file, "%s: %s: (file%s:  function:%s line:%d) %s",
			logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
		fmt.Fprintln(file)
	}
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
	LogData := writeLog(LeDeBug, fmtstr, args...)
	select {
	case f.logDataChan <- LogData:
	default:
	}
}

func (f *Filelog) Trace(fmtstr string, args ...interface{}) {
	if f.level > LeTrace {
		return
	}
	LogData := writeLog(LeTrace, fmtstr, args...)
	select {
	case f.logDataChan <- LogData:
	default:
	}
}

func (f *Filelog) Info(fmtstr string, args ...interface{}) {
	if f.level > LeInfo {
		return
	}
	LogData := writeLog(LeInfo, fmtstr, args...)
	select {
	case f.logDataChan <- LogData:
	default:
	}
}

func (f *Filelog) Warn(fmtstr string, args ...interface{}) {
	if f.level > LeWarn {
		return
	}
	LogData := writeLog(LeWarn, fmtstr, args...)
	select {
	case f.logDataChan <- LogData:
	default:
	}
}

func (f *Filelog) Error(fmtstr string, args ...interface{}) {
	if f.level > LeError {
		return
	}
	LogData := writeLog(LeError, fmtstr, args...)
	select {
	case f.logDataChan <- LogData:
	default:
	}
}

func (f *Filelog) Fatal(fmtstr string, args ...interface{}) {
	LogData := writeLog(LeFatal, fmtstr, args...)
	select {
	case f.logDataChan <- LogData:
	default:
	}
}

func (f *Filelog) Close() {
	f.file.Close()
	f.warnFile.Close()
}
