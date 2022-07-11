# bvmgo-util/errors

## Create error

Creation functions:
* `errors.NewWithCause(cause error, format string, attributes ...interface{}) error` build a new error with a cause and a message.
* `errors.New(format string, attributes ...interface{}) error` build a new error with a message.
* `errors.Wrap(cause error) error` build a new error with the cause (only add the stack information).

Creation functions return `error` type.
Errors implement `errors.TraceableError` interface.

## `errors.TraceableError` interface

`TraceableError` is an error with cause and error position information.
