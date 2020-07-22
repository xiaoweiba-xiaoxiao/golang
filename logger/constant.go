package logger

const (
	LeDeBug = iota
	LeTrace
	LeInfo
	LeWarn
	LeError
	LeFatal
)

func getLevel(level int) string {
	switch level {
	case LeDeBug:
		return "DEBUG"
	case LeTrace:
		return "TRACE"
	case LeInfo:
		return "INFO"
	case LeWarn:
		return "WARN"
	case LeError:
		return "ERROR"
	case LeFatal:
		return "FATAL"
	default:
		return "DEBUG"
	}
}
