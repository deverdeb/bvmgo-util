package logs

import (
	"log"
)

// DefaultFormater is the default log formatter
var DefaultFormater Formatter = &defaultFormaterImpl{}

// DefaultLogLevel is the default log level
var DefaultLogLevel = LevelInfo

var DefaultLogger = New("")

func init() {
	log.Default().SetFlags(0)
	log.Default().SetPrefix("")
}

func Trace(attributes ...any) {
	DefaultLogger.log(LevelTrace, attributes...)
}
func Tracef(format string, attributes ...any) {
	DefaultLogger.logf(LevelTrace, format, attributes...)
}
func Debug(attributes ...any) {
	DefaultLogger.log(LevelDebug, attributes...)
}
func Debugf(format string, attributes ...any) {
	DefaultLogger.logf(LevelDebug, format, attributes...)
}
func Info(attributes ...any) {
	DefaultLogger.log(LevelInfo, attributes...)
}
func Infof(format string, attributes ...any) {
	DefaultLogger.logf(LevelInfo, format, attributes...)
}
func Warn(attributes ...any) {
	DefaultLogger.log(LevelWarn, attributes...)
}
func Warnf(format string, attributes ...any) {
	DefaultLogger.logf(LevelWarn, format, attributes...)
}
func Error(attributes ...any) {
	DefaultLogger.log(LevelError, attributes...)
}
func Errorf(format string, attributes ...any) {
	DefaultLogger.logf(LevelError, format, attributes...)
}
func Fatal(attributes ...any) {
	DefaultLogger.log(LevelFatal, attributes...)
}
func Fatalf(format string, attributes ...any) {
	DefaultLogger.logf(LevelFatal, format, attributes...)
}
