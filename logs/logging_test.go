package logs

import (
	"os"
)

func ExampleLogger_Trace() {
	mockBeforeTest()
	defer restoreAfterTest()
	logger := New("ExampleLoggerTrace")
	logger.Output().SetOutput(os.Stdout)

	logger.SetLevel(LevelTrace)
	logger.Trace("hello everybody")
	logger.Tracef("hello %d", 123456)
	logger.SetLevel(LevelDebug)
	logger.Trace("hidden log")
	logger.SetLevel(LevelInfo)
	logger.Tracef("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [TRACE] ExampleLoggerTrace - hello everybody ( logging_test.go:14 )
	// 1982-03-15T12:56:14 [TRACE] ExampleLoggerTrace - hello 123456 ( logging_test.go:15 )
}

func ExampleLogger_Debug() {
	mockBeforeTest()
	defer restoreAfterTest()
	logger := New("ExampleLoggerDebug")
	logger.Output().SetOutput(os.Stdout)

	logger.SetLevel(LevelTrace)
	logger.Debug("hello everybody")
	logger.SetLevel(LevelDebug)
	logger.Debugf("hello %s", "world")
	logger.SetLevel(LevelInfo)
	logger.Debug("hidden log")
	logger.SetLevel(LevelWarn)
	logger.Debugf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [DEBUG] ExampleLoggerDebug - hello everybody ( logging_test.go:33 )
	// 1982-03-15T12:56:14 [DEBUG] ExampleLoggerDebug - hello world ( logging_test.go:35 )
}

func ExampleLogger_Info() {
	mockBeforeTest()
	defer restoreAfterTest()
	logger := New("ExampleLoggerInfo")
	logger.Output().SetOutput(os.Stdout)

	logger.SetLevel(LevelTrace)
	logger.Info("hello everybody")
	logger.SetLevel(LevelInfo)
	logger.Infof("hello %v", false)
	logger.SetLevel(LevelWarn)
	logger.Info("hidden log")
	logger.SetLevel(LevelError)
	logger.Infof("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ INFO] ExampleLoggerInfo - hello everybody ( logging_test.go:53 )
	// 1982-03-15T12:56:14 [ INFO] ExampleLoggerInfo - hello false ( logging_test.go:55 )
}

func ExampleLogger_Warn() {
	mockBeforeTest()
	defer restoreAfterTest()
	logger := New("ExampleLoggerWarn")
	logger.Output().SetOutput(os.Stdout)

	logger.SetLevel(LevelInfo)
	logger.Warn("hello everybody")
	logger.SetLevel(LevelWarn)
	logger.Warnf("hello %v", "world")
	logger.SetLevel(LevelError)
	logger.Warn("hidden log")
	logger.SetLevel(LevelFatal)
	logger.Warnf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ WARN] ExampleLoggerWarn - hello everybody ( logging_test.go:73 )
	// 1982-03-15T12:56:14 [ WARN] ExampleLoggerWarn - hello world ( logging_test.go:75 )
}

func ExampleLogger_Error() {
	mockBeforeTest()
	defer restoreAfterTest()
	logger := New("ExampleLoggerError")
	logger.Output().SetOutput(os.Stdout)

	logger.SetLevel(LevelInfo)
	logger.Error("hello everybody")
	logger.SetLevel(LevelError)
	logger.Errorf("hello %v", "world")
	logger.SetLevel(LevelFatal)
	logger.Error("hidden log")
	logger.SetLevel(LevelFatal)
	logger.Errorf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ERROR] ExampleLoggerError - hello everybody ( logging_test.go:93 )
	// 1982-03-15T12:56:14 [ERROR] ExampleLoggerError - hello world ( logging_test.go:95 )
}

func ExampleLogger_Fatal() {
	mockBeforeTest()
	defer restoreAfterTest()
	logger := New("ExampleLoggerFatal")
	logger.Output().SetOutput(os.Stdout)

	logger.SetLevel(LevelError)
	logger.Fatal("hello everybody")
	logger.SetLevel(LevelFatal)
	logger.Fatalf("hello %v", "world")

	// Output:
	// 1982-03-15T12:56:14 [FATAL] ExampleLoggerFatal - hello everybody ( logging_test.go:113 )
	// exit with code 1
	// 1982-03-15T12:56:14 [FATAL] ExampleLoggerFatal - hello world ( logging_test.go:115 )
	// exit with code 1
}
