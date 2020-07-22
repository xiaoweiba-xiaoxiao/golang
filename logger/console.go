package logger

import (
	"os"
)

type Console struct {
	level int
}

func NewConsole(level int) Logger {
	return &Console{
		level: level,
	}
}

func (c *Console) Setlevel(level int) {
	c.level = level
}
func (c *Console) DeBug(fmtstr string, args ...interface{}) {
	if c.level > LeDeBug {
		return
	}
	writeLog(os.Stdout, LeDeBug, fmtstr, args...)
}
func (c *Console) Trace(fmtstr string, args ...interface{}) {
	if c.level > LeTrace {
		return
	}
	writeLog(os.Stdout, LeTrace, fmtstr, args...)
}
func (c *Console) Info(fmtstr string, args ...interface{}) {
	if c.level > LeInfo {
		return
	}
	writeLog(os.Stdout, LeInfo, fmtstr, args...)
}
func (c *Console) Warn(fmtstr string, args ...interface{}) {
	if c.level > LeWarn {
		return
	}
	writeLog(os.Stdout, LeWarn, fmtstr, args...)
}
func (c *Console) Error(fmtstr string, args ...interface{}) {
	if c.level > LeError {
		return
	}
	writeLog(os.Stdout, LeError, fmtstr, args...)
}
func (c *Console) Fatal(fmtstr string, args ...interface{}) {
	writeLog(os.Stdout, LeFatal, fmtstr, args...)
}
func (c *Console) Close() {

}
