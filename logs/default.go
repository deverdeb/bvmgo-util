package logs

import (
	"log"
	"os"
)

// DefaultFormater is the default message formatter for loggers.
var DefaultFormater Formatter = &defaultFormaterImpl{}

// DefaultLogLevel is the default log level for loggers.
var DefaultLogLevel = LevelInfo

// DefaultOuput is the default output for loggers.
var DefaultOuput = log.New(os.Stdout, "", 0)

// defaultLogger is the default logger.
var defaultLogger = New("")

// DefaultLogger returns the default logger.
func DefaultLogger() *Logger {
	return defaultLogger
}

func Trace(attributes ...any) {
	defaultLogger.log(LevelTrace, attributes...)
}
func Tracef(format string, attributes ...any) {
	defaultLogger.logf(LevelTrace, format, attributes...)
}
func Debug(attributes ...any) {
	defaultLogger.log(LevelDebug, attributes...)
}
func Debugf(format string, attributes ...any) {
	defaultLogger.logf(LevelDebug, format, attributes...)
}
func Info(attributes ...any) {
	defaultLogger.log(LevelInfo, attributes...)
}
func Infof(format string, attributes ...any) {
	defaultLogger.logf(LevelInfo, format, attributes...)
}
func Warn(attributes ...any) {
	defaultLogger.log(LevelWarn, attributes...)
}
func Warnf(format string, attributes ...any) {
	defaultLogger.logf(LevelWarn, format, attributes...)
}
func Error(attributes ...any) {
	defaultLogger.log(LevelError, attributes...)
}
func Errorf(format string, attributes ...any) {
	defaultLogger.logf(LevelError, format, attributes...)
}
func Fatal(attributes ...any) {
	defaultLogger.log(LevelFatal, attributes...)
	exit(1)
}
func Fatalf(format string, attributes ...any) {
	defaultLogger.logf(LevelFatal, format, attributes...)
	exit(1)
}
