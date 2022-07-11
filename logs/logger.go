package logs

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// now function returns current time.
var now = time.Now

// exit function exits application with the given status code.
var exit = os.Exit

// Logger is a logger...
type Logger struct {
	// prefix is the logger message prefix
	prefix string
	// level is the minimal level to log message
	level LogLevel
	// formatter is the log message formatter
	formatter Formatter
	// output is the output destination for the logger
	output *log.Logger
}

// New method creates a new logger.
//
// prefix parameter define logger prefix.
// Example:
//
//   logger1 := logs.New("MyLogger") // logger with prefix
//   logger2 := logs.New("") // logger without prefix
func New(prefix string) *Logger {
	return &Logger{
		prefix:    prefix,
		formatter: DefaultFormater,
		level:     LevelDefault,
		output:    DefaultOuput,
	}
}

// Prefix returns the output prefix for the logger.
func (logger *Logger) Prefix() string {
	return logger.prefix
}

// Level returns the log Level for the logger.
func (logger *Logger) Level() LogLevel {
	if logger.level == LevelDefault {
		return DefaultLogLevel
	}
	return logger.level
}

// SetLevel sets the log Level for the logger.
func (logger *Logger) SetLevel(level LogLevel) {
	logger.level = level
}

// Formatter returns the message Formatter for the logger.
func (logger *Logger) Formatter() Formatter {
	return logger.formatter
}

// SetFormatter sets the message Formatter for the logger.
func (logger *Logger) SetFormatter(formatter Formatter) {
	logger.formatter = formatter
}

// Output returns the output destination for the logger.
func (logger *Logger) Output() *log.Logger {
	return logger.output
}

// SetOutput sets the destination for the logger.
func (logger *Logger) SetOutput(output *log.Logger) {
	logger.output = output
}

// Trace writes message with "Trace" level.
func (logger *Logger) Trace(attributes ...any) {
	logger.log(LevelTrace, attributes...)
}

// Tracef writes message with "Trace" level.
func (logger *Logger) Tracef(format string, attributes ...any) {
	logger.logf(LevelTrace, format, attributes...)
}

// Debug writes message with "Debug" level.
func (logger *Logger) Debug(attributes ...any) {
	logger.log(LevelDebug, attributes...)
}

// Debugf writes message with "Debug" level.
func (logger *Logger) Debugf(format string, attributes ...any) {
	logger.logf(LevelDebug, format, attributes...)
}

// Info writes message with "Info" level.
func (logger *Logger) Info(attributes ...any) {
	logger.log(LevelInfo, attributes...)
}

// Infof writes message with "Info" level.
func (logger *Logger) Infof(format string, attributes ...any) {
	logger.logf(LevelInfo, format, attributes...)
}

// Warn writes message with "Warn" level.
func (logger *Logger) Warn(attributes ...any) {
	logger.log(LevelWarn, attributes...)
}

// Warnf writes message with "Warn" level.
func (logger *Logger) Warnf(format string, attributes ...any) {
	logger.logf(LevelWarn, format, attributes...)
}

// Error writes message with "Error" level.
func (logger *Logger) Error(attributes ...any) {
	logger.log(LevelError, attributes...)
}

// Errorf writes message with "Error" level.
func (logger *Logger) Errorf(format string, attributes ...any) {
	logger.logf(LevelError, format, attributes...)
}

// Fatal writes message with "Fatal" level.
// And call os.Exit(1).
func (logger *Logger) Fatal(attributes ...any) {
	logger.log(LevelFatal, attributes...)
	exit(1)
}

// Fatalf writes message with "Fatal" level.
// And call os.Exit(1).
func (logger *Logger) Fatalf(format string, attributes ...any) {
	logger.logf(LevelFatal, format, attributes...)
	exit(1)
}

// log checks level, formats message and writes to output.
func (logger *Logger) log(level LogLevel, attributes ...any) {
	if level >= logger.Level() {
		file, line := extractFilenameAndLine(2)
		args, err := extractErrorOfArguments(attributes...)
		message := fmt.Sprint(args...)
		logger.output.Print(logger.formatter.Format(logger.prefix, level, file, line, now(), err, message))
	}
}

// logf checks level, formats message and writes to output.
func (logger *Logger) logf(level LogLevel, format string, attributes ...any) {
	if level >= logger.level {
		file, line := extractFilenameAndLine(2)
		args, err := extractErrorOfArguments(attributes...)
		message := fmt.Sprintf(format, args...)
		logger.output.Print(logger.formatter.Format(logger.prefix, level, file, line, now(), err, message))
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
