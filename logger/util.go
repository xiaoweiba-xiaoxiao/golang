package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

func writeLog(file *os.File, level int, fmtstr string, args ...interface{}) {
	// if l.Setlevel(level) > level {
	// 	return
	// }
	timeStr := time.Now().Format("2006-01-02 15:04:05.999")
	levelStr := getLevel(level)
	fileInfo, funcInfo, lineInfo := GetLineInfo()
	fileInfo = path.Base(fileInfo)
	funcInfo = path.Base(funcInfo)
	msg := fmt.Sprintf(fmtstr, args...)
	fmt.Fprintf(file, "%s: %s: (file%s:  function:%s line:%d) %s", timeStr, levelStr, fileInfo, funcInfo, lineInfo, msg)
	fmt.Fprintln(file)
}
