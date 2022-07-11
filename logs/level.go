package logs

type LogLevel uint8

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelDefault
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (level LogLevel) String() string {
	switch level {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelDefault, LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "?????"
	}
}
