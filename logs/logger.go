package logs

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// now function return current time
var now = time.Now

// Logger is a logger...
type Logger struct {
	// prefix is the logger message prefix
	prefix string
	// level is the minimal level to log message
	level LogLevel
	// formatter is the log message formatter
	formatter Formatter
}

// New method creates a new logger
func New(prefix string) *Logger {
	return &Logger{
		prefix:    prefix,
		formatter: DefaultFormater,
		level:     DefaultLogLevel,
	}
}

func (logger *Logger) Prefix() string {
	return logger.prefix
}

func (logger *Logger) SetPrefix(prefix string) {
	logger.prefix = prefix
}

func (logger *Logger) Level() LogLevel {
	return logger.level
}

func (logger *Logger) SetLevel(level LogLevel) {
	logger.level = level
}

func (logger *Logger) Formatter() Formatter {
	return logger.formatter
}

func (logger *Logger) SetFormatter(formatter Formatter) {
	logger.formatter = formatter
}

func (logger *Logger) Trace(attributes ...any) {
	logger.log(LevelTrace, attributes...)
}
func (logger *Logger) Tracef(format string, attributes ...any) {
	logger.logf(LevelTrace, format, attributes...)
}
func (logger *Logger) Debug(attributes ...any) {
	logger.log(LevelDebug, attributes...)
}
func (logger *Logger) Debugf(format string, attributes ...any) {
	logger.logf(LevelDebug, format, attributes...)
}
func (logger *Logger) Info(attributes ...any) {
	logger.log(LevelInfo, attributes...)
}
func (logger *Logger) Infof(format string, attributes ...any) {
	logger.logf(LevelInfo, format, attributes...)
}
func (logger *Logger) Warn(attributes ...any) {
	logger.log(LevelWarn, attributes...)
}
func (logger *Logger) Warnf(format string, attributes ...any) {
	logger.logf(LevelWarn, format, attributes...)
}
func (logger *Logger) Error(attributes ...any) {
	logger.log(LevelError, attributes...)
}
func (logger *Logger) Errorf(format string, attributes ...any) {
	logger.logf(LevelError, format, attributes...)
}
func (logger *Logger) Fatal(attributes ...any) {
	logger.log(LevelFatal, attributes...)
}
func (logger *Logger) Fatalf(format string, attributes ...any) {
	logger.logf(LevelFatal, format, attributes...)
}

func (logger *Logger) log(level LogLevel, attributes ...any) {
	if level >= logger.level {
		file, line := extractFilenameAndLine(2)
		args, err := extractErrorOfArguments(attributes...)
		message := fmt.Sprint(args...)
		log.Print(logger.formatter.Format(logger.prefix, level, file, line, now(), err, message))
	}
}
func (logger *Logger) logf(level LogLevel, format string, attributes ...any) {
	if level >= logger.level {
		file, line := extractFilenameAndLine(2)
		args, err := extractErrorOfArguments(attributes...)
		message := fmt.Sprintf(format, args...)
		log.Print(logger.formatter.Format(logger.prefix, level, file, line, now(), err, message))
	}
}

// extractError extract last attribut if it is an error (error type).
// If last attributs is an error, return attributes without last error and error
// Else, return all attributes and nil for error
func (logger *Logger) extractError(attributes ...interface{}) ([]interface{}, error) {
	nbAttributes := len(attributes)
	if nbAttributes > 0 {
		lastAttribute, isError := attributes[nbAttributes-1].(error)
		if isError {
			return attributes[0 : nbAttributes-1], lastAttribute
		} else {
			return attributes, nil
		}
	} else {
		return nil, nil
	}
}

// extractFilenameAndLine returns the log position in code (filename and line).
// The argument skip is the number of stack frames to ascend, with 0 identifying the caller of extractFilenameAndLine.
func extractFilenameAndLine(skip int) (filename string, line int) {
	// Extract element - skip +1 -> ignore extractFilenameAndLine(...) function
	_, file, line, ok := runtime.Caller(skip + 1)
	if ok {
		return file, line
	} else {
		// Error to extract execution stack.
		// Return unknown file at line 0
		return "<unknown>", 0
	}
}

// extractErrorOfArguments returns the last arguments if last argument is an error.
// Return nil if last argument is not an error or if attributes array is empty.
func extractErrorOfArguments(attributes ...interface{}) ([]interface{}, error) {
	if attributes == nil || len(attributes) == 0 {
		return attributes, nil
	}
	lastArg := attributes[len(attributes)-1]
	err, ok := lastArg.(error)
	if ok {
		return attributes[0 : len(attributes)-1], err
	} else {
		return attributes, nil
	}
}
