package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Filelog struct {
	level        int
	logPath      string
	logFile      string
	file         *os.File
	warnFile     *os.File
	logDataChan  chan *LogData
	logSpiltType string
	logSplitSize int64
	lastDay      int
}

func NewFileLog(msg map[string]string) (log Logger, err error) {
	var filelog *Filelog = &Filelog{}
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
	logSpiltType, ok := msg["log_split_type"]
	if !ok {
		logSpiltType = "time"
	}
	if logSpiltType == "size" {
		logSplitSizeStr, ok := msg["log_split_size"]
		if !ok {
			logSplitSizeStr = "104857600"
		}
		logSplitSize, err := strconv.ParseInt(logSplitSizeStr, 10, 64)
		if err != nil {
			logSplitSize = 104857600
		}
		fmt.Println(logSplitSize)
		filelog.logSplitSize = logSplitSize
	} else {
		logSpiltType = "time"
	}
	fmt.Println(logSpiltType)
	filelog.level = level
	filelog.logFile = LogFile
	filelog.logPath = LogPath
	filelog.logDataChan = make(chan *LogData, chanNum)
	filelog.logSpiltType = logSpiltType
	filelog.lastDay = time.Now().Day()
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

func (f *Filelog) splitByTime(warnTag bool) {
	now := time.Now()
	day := now.Day()
	if day == f.lastDay {
		return
	}
	file := f.file
	newLogFilestr := fmt.Sprintf("%s/%s_%04d%02d%02d", f.logPath, f.logFile, now.Year(), now.Month(), now.Day())
	filename := fmt.Sprintf("%s/%s", f.logPath, f.logFile)
	if warnTag {
		newLogFilestr = fmt.Sprintf("%s/%s.wf_%04d%02d%02d", f.logPath, f.logFile, now.Year(), now.Month(), now.Day())
		filename = fmt.Sprintf("%s/%s.wf", f.logPath, f.logFile)
		file = f.warnFile
	}
	file.Close()
	os.Rename(filename, newLogFilestr)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return
	}
	if warnTag {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *Filelog) splitBySize(warnTag bool) {
	file := f.file
	if warnTag {
		file = f.warnFile
	}
	infoStr, err := file.Stat()
	if err != nil {
		return
	}
	if fileSize := infoStr.Size(); fileSize < f.logSplitSize {
		return
	}
	now := time.Now()
	newLogFilestr := fmt.Sprintf("%s/%s_%04d%02d%02d%02d%02d%02d", f.logPath, f.logFile, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	filename := fmt.Sprintf("%s/%s", f.logPath, f.logFile)
	if warnTag {
		newLogFilestr = fmt.Sprintf("%s/%s.wf_%04d%02d%02d%02d%02d%02d", f.logPath, f.logFile, now.Year(), now.Month(), now.Day())
		filename = fmt.Sprintf("%s/%s.wf", f.logPath, f.logFile)
	}
	file.Close()
	os.Rename(filename, newLogFilestr)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return
	}
	if warnTag {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *Filelog) checkSpilt(warnTag bool) {
	switch f.logSpiltType {
	case "time":
		f.splitByTime(warnTag)
	case "size":
		f.splitBySize(warnTag)
	}
}

func (f *Filelog) writelog_d() {
	for logData := range f.logDataChan {
		var file = f.file
		if logData.WarnOrFatal {
			file = f.warnFile
		}
		f.checkSpilt(logData.WarnOrFatal)
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
