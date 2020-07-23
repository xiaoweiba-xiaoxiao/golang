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

func stringtoLevel(level string) int{
	switch level {
	case "debug":
		return LeDeBug
	case "trace":
		return LeTrace
	case "info":
		return LeInfo
	case "warn":
		return LeWarn
	case "error":
		return LeError
	case "fatal":
		return LeFatal
	}
	return LeDeBug
}
