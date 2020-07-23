package logger

import (
	"fmt"
	"os"
)

type Console struct {
	level int
}

func NewConsole(config map[string]string) (log Logger, err error) {
	level, ok := config["level"]
	if !ok {
		level = "debug"
	}
	Level := stringtoLevel(level)
	if err != nil {
		Level = 3
	}
	log = &Console{
		level: Level,
	}
	return
}

func (c *Console) Setlevel(level int) {
	c.level = level
}
func (c *Console) DeBug(fmtstr string, args ...interface{}) {
	if c.level > LeDeBug {
		return
	}
	logData := writeLog(LeDeBug, fmtstr, args...)
	fmt.Fprintf(os.Stdout, "%s: %s: (file%s:  function:%s line:%d) %s", logData.TimeStr,
		logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *Console) Trace(fmtstr string, args ...interface{}) {
	if c.level > LeTrace {
		return
	}
	logData := writeLog(LeTrace, fmtstr, args...)
	fmt.Fprintf(os.Stdout, "%s: %s: (file%s:  function:%s line:%d) %s",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *Console) Info(fmtstr string, args ...interface{}) {
	if c.level > LeInfo {
		return
	}
	logData := writeLog(LeInfo, fmtstr, args...)
	fmt.Fprintf(os.Stdout, "%s: %s: (file%s:  function:%s line:%d) %s",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *Console) Warn(fmtstr string, args ...interface{}) {
	if c.level > LeWarn {
		return
	}
	logData := writeLog(LeWarn, fmtstr, args...)
	fmt.Fprintf(os.Stdout, "%s: %s: (file%s:  function:%s line:%d) %s",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *Console) Error(fmtstr string, args ...interface{}) {
	if c.level > LeError {
		return
	}
	logData := writeLog(LeError, fmtstr, args...)
	fmt.Fprintf(os.Stdout, "%s: %s: (file%s:  function:%s line:%d) %s",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *Console) Fatal(fmtstr string, args ...interface{}) {
	logData := writeLog(LeFatal, fmtstr, args...)
	fmt.Fprintf(os.Stdout, "%s: %s: (file%s:  function:%s line:%d) %s",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *Console) Close() {

}
