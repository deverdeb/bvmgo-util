package errors

import (
	"fmt"
	"runtime"
)

// unknownFunctionLabel is the label for unknown functions.
var unknownFunctionLabel = "<unknown>"

// TraceableError is an error with cause and stacktrace.
type TraceableError interface {
	// Message method returns error message
	Message() string
	// File method returns error filename
	File() string
	// Line method returns line number of error in file
	Line() int
	// Cause method returns cause error
	Cause() error
}

// customError is an error with cause.
type customError struct {
	// message is the error label
	message string
	// cause is the wrapped error
	cause error
	// file is the filename
	file string
	// function is the name of function in file
	function string
	// line is the line in file
	line int
}

// NewWithCause build a new error with a cause and a message.
func NewWithCause(cause error, format string, attributes ...interface{}) error {
	return extractPositionAndBuildCustomError(cause, 1, format, attributes...)
}

// New build a new error with a message.
func New(format string, attributes ...interface{}) error {
	return extractPositionAndBuildCustomError(nil, 1, format, attributes...)
}

// Wrap build a new error with the cause (only add the stack information).
func Wrap(cause error) error {
	err, ok := cause.(TraceableError)
	if ok {
		return extractPositionAndBuildCustomError(cause, 1, err.Message())
	} else {
		return extractPositionAndBuildCustomError(cause, 1, cause.Error())
	}

}

// extractPositionAndBuildCustomError return a new error.
//
// The argument cause is the next error in  the error chain (see Wrapper error interface). It is nullable.
// The argument skip is the number of stack frames to ascend, with 0 identifying the caller.
// The argument format is the format of error message (see fmt.Sprintf method format).
// The argument attributes are the arguments to complete the error message (see fmt.Sprintf method format).
func extractPositionAndBuildCustomError(cause error, execStackSkip int, format string, attributes ...interface{}) error {
	file, function, line := extractPositionInExecutionStack(execStackSkip + 1) // +1 to skip current method.
	return buildCustomError(cause, file, function, line, format, attributes...)
}

// buildCustomError return a new error.
func buildCustomError(cause error, file string, function string, line int, format string, attributes ...interface{}) error {
	return &customError{
		message:  fmt.Sprintf(format, attributes...),
		cause:    cause,
		file:     file,
		function: function,
		line:     line,
	}
}

// Error returns the error message.
// If error has cause, add the cause message.
func (err *customError) Error() (message string) {
	cause := ""
	if err.cause != nil {
		cause = "\n    > cause by: " + err.cause.Error()
	}
	position := ""
	if err.line > 0 {
		if err.function == unknownFunctionLabel {
			position = fmt.Sprintf(" ( at %s:%d )", err.file, err.line)
		} else {
			position = fmt.Sprintf(" ( at %s:%d )", err.function, err.line)
		}
	}
	return err.message + position + cause
}

// Message returns only the error message.
// If error has cause, add the cause message.
func (err *customError) Message() (message string) {
	return err.message
}

// Unwrap method returns cause error.
func (err *customError) Unwrap() error {
	return err.cause
}

// File method returns error filename
func (err *customError) File() string {
	return err.file
}

// Function method returns function of error in file
func (err *customError) Function() string {
	return err.function
}

// Line method returns line number of error in file
func (err *customError) Line() int {
	return err.line
}

// Cause method returns cause error.
func (err *customError) Cause() error {
	return err.cause
}

// extractPositionInExecutionStack returns the execution position (file, function and line).
// The argument skip is the number of stack frames to ascend, with 0 identifying the caller of extractPositionInExecutionStack.
func extractPositionInExecutionStack(skip int) (file string, function string, line int) {
	// Extract element - skip +1 -> ignore extractPositionInExecutionStack(...)
	ptrExec, file, line, ok := runtime.Caller(skip + 1)
	if ok {
		execFrames := runtime.CallersFrames([]uintptr{ptrExec})
		if execFrames != nil {
			frame, _ := execFrames.Next()
			return frame.File, frame.Function, frame.Line
		} else {
			// No frame... return unknown function
			return file, unknownFunctionLabel, line
		}
	} else {
		// Error to extract execution stack.
		// Return default runtime.Caller(...) values and unknown function
		return file, unknownFunctionLabel, line
	}
}
