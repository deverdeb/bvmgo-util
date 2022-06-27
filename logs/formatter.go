package logs

import (
	goerr "errors"
	"fmt"
	"github.com/deverdeb/bvmgo-util/errors"
	"path/filepath"
	"time"
)

// Formatter is an interface to format log messages.
type Formatter interface {
	// Format return a formatted log message
	Format(prefix string, level LogLevel, fileName string, fileLine int,
		date time.Time, err error, message string) string
}

// defaultFormaterImpl is a basic Formatter implementation.
type defaultFormaterImpl struct {
}

// Format return a formatted log message
func (formatter *defaultFormaterImpl) Format(prefix string, level LogLevel, fileName string, fileLine int,
	date time.Time, err error, message string) string {
	strDate := date.Format("2006-01-02T15:04:05")
	result := fmt.Sprintf("%s [%5s] ", strDate, level.String())
	if prefix != "" {
		result += prefix + " - "
	}
	result += message
	if fileLine > 0 {
		result += fmt.Sprintf(" ( %s:%d )", filepath.Base(fileName), fileLine)
	}
	if err != nil {
		result += "\n  > error: " + FormatError(err, -1)
	}
	return result
}

// FormatError converts an error to log message
func FormatError(err error, errorsDepth int) string {
	if err == nil {
		return "nil"
	}
	if errorsDepth == 0 {
		return "..."
	}
	var result string
	traceableError, ok := err.(errors.TraceableError)
	if ok {
		result = fmt.Sprintf("%s ( %s:%d )", traceableError.Message(), filepath.Base(traceableError.File()), traceableError.Line())
	} else {
		result = err.Error()
	}
	cause := goerr.Unwrap(err)
	if cause != nil {
		result += "\n  > cause by: " + FormatError(cause, errorsDepth-1)
	}
	return result
}
