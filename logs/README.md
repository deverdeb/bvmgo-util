# bvmgo-util/logs

## Loggers

### Create logger

`func logs.New(prefix string) *Logger` function creates a new logger.

```go
// With prefix
logger := logs.New("Logger prefix")
```

```go
// Without prefix
logger := logs.New("")
```

### Log messages

Logs can have 6 levels : Trace, Debug, Info, Warn, Error and Fatal.

Note : Fatal log call `os.exit(1)` function.

We can log message with logger functions :
* `logs.Trace(attributes ...any)` or `logs.Tracef(format string, attributes ...any)`
* `logs.Debug(attributes ...any)` or `logs.Debugf(format string, attributes ...any)`
* `logs.Info(attributes ...any)` or `logs.Infof(format string, attributes ...any)`
* `logs.Warn(attributes ...any)` or `logs.Warnf(format string, attributes ...any)`
* `logs.Error(attributes ...any)` or `logs.Errorf(format string, attributes ...any)`
* `logs.Fatal(attributes ...any)` or `logs.Fatalf(format string, attributes ...any)`

## Default logger

We can use default logger with functions :
* `logs.Trace(attributes ...any)` or `logs.Tracef(format string, attributes ...any)`
* `logs.Debug(attributes ...any)` or `logs.Debugf(format string, attributes ...any)`
* `logs.Info(attributes ...any)` or `logs.Infof(format string, attributes ...any)`
* `logs.Warn(attributes ...any)` or `logs.Warnf(format string, attributes ...any)`
* `logs.Error(attributes ...any)` or `logs.Errorf(format string, attributes ...any)`
* `logs.Fatal(attributes ...any)` or `logs.Fatalf(format string, attributes ...any)`

`logs.DefaultLogger() *Logger` function returns the default logger.

Default logger has not prefix. 

## Configuration

### Logger prefix

Logger prefix is define when logger is created.
Prefix cannot be redefined.

### Logger level

`logs.DefaultLogLevel` define the default log level.

By default, logger use the default log level.
But we can change logger level.
* `logs.Logger.Level() LogLevel` returns logger level.
* `logs.Logger.SetLevel(level LogLevel)` sets logger level.

Log level can have 6 values : `LevelTrace`, `LevelDebug`, `LevelInfo`, `LevelWarn`, `LevelError` and `LevelFatal`.

`LevelDefault` special level replace logger level by value of default `logs.DefaultLogLevel`.

### Logger formatter

Log message formatter can be redefined by a `Formatter` interface implementation. 

We can change logger formatter.
* `logs.Logger.Formatter() Formatter` returns logger formatter.
* `logs.Logger.SetFormatter(formatter Formatter)` sets logger formatter.

`logs.FormatError(err error, errorsDepth int) string` function can be used to format error with wrapped errors.

### Logger output

Logger use default golang `log.Logger` to log messages.

We can customize output logger with functions:
* `logs.Logger.Output() *log.Logger` returns logger output.
* `logs.Logger.SetOutput(output *log.Logger)` sets logger output.
