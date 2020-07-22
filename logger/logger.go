package logger

type Logger interface {
	Setlevel(level int)
	DeBug(fmtstr string, args ...interface{})
	Trace(fmtstr string, args ...interface{})
	Info(fmtstr string, args ...interface{})
	Warn(fmtstr string, args ...interface{})
	Error(fmtstr string, args ...interface{})
	Fatal(fmtstr string, args ...interface{})
	Close()
}
