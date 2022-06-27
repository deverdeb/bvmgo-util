package logs // FIXME A virer

import (
	"fmt"
	"github.com/deverdeb/bvmgo-util/errors"
	"io"
	"log"
	"os"
	"time"
)

var previousWriter io.Writer

// mockBeforeTest replace "Now()" function by a mock and log writer.
// Mocked "Now()" function returns 2018-12-15T17:8:12.1565 time.
func mockBeforeTest() {
	previousWriter = log.Writer()
	log.SetOutput(os.Stdout)
	now = func() time.Time {
		return time.Date(1982, time.March, 15, 12, 56, 14, 123, time.Local)
	}
}

// restoreAfterTest restore the default "Now()" function and log writer.
func restoreAfterTest() {
	log.SetOutput(previousWriter)
	now = time.Now
}

func ExampleTrace() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger.SetLevel(LevelTrace)
	Trace("hello everybody")
	Tracef("hello %d", 123456)
	DefaultLogger.SetLevel(LevelDebug)
	Trace("hidden log")
	DefaultLogger.SetLevel(LevelInfo)
	Tracef("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [TRACE] hello everybody ( logging_test.go:35 )
	// 1982-03-15T12:56:14 [TRACE] hello 123456 ( logging_test.go:36 )
}

func ExampleDebug() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger.SetLevel(LevelTrace)
	Debug("hello everybody")
	DefaultLogger.SetLevel(LevelDebug)
	Debugf("hello %s", "world")
	DefaultLogger.SetLevel(LevelInfo)
	Debug("hidden log")
	DefaultLogger.SetLevel(LevelWarn)
	Debugf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [DEBUG] hello everybody ( logging_test.go:52 )
	// 1982-03-15T12:56:14 [DEBUG] hello world ( logging_test.go:54 )
}

func ExampleInfo() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger.SetLevel(LevelTrace)
	Info("hello everybody")
	DefaultLogger.SetLevel(LevelInfo)
	Infof("hello %v", false)
	DefaultLogger.SetLevel(LevelWarn)
	Info("hidden log")
	DefaultLogger.SetLevel(LevelError)
	Infof("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ INFO] hello everybody ( logging_test.go:70 )
	// 1982-03-15T12:56:14 [ INFO] hello false ( logging_test.go:72 )
}

func ExampleWarn() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger.SetLevel(LevelInfo)
	Warn("hello everybody")
	DefaultLogger.SetLevel(LevelWarn)
	Warnf("hello %v", "world")
	DefaultLogger.SetLevel(LevelError)
	Warn("hidden log")
	DefaultLogger.SetLevel(LevelFatal)
	Warnf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ WARN] hello everybody ( logging_test.go:88 )
	// 1982-03-15T12:56:14 [ WARN] hello world ( logging_test.go:90 )
}

func ExampleError() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger.SetLevel(LevelInfo)
	Error("hello everybody")
	DefaultLogger.SetLevel(LevelError)
	Errorf("hello %v", "world")
	DefaultLogger.SetLevel(LevelFatal)
	Error("hidden log")
	DefaultLogger.SetLevel(LevelFatal)
	Errorf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ERROR] hello everybody ( logging_test.go:106 )
	// 1982-03-15T12:56:14 [ERROR] hello world ( logging_test.go:108 )
}

func ExampleFatal() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger.SetLevel(LevelError)
	Fatal("hello everybody")
	DefaultLogger.SetLevel(LevelFatal)
	Fatalf("hello %v", "world")

	// Output:
	// 1982-03-15T12:56:14 [FATAL] hello everybody ( logging_test.go:124 )
	// 1982-03-15T12:56:14 [FATAL] hello world ( logging_test.go:126 )
}

func ExampleLogger_Formatter() {
	mockBeforeTest()
	defer restoreAfterTest()

	err1 := fmt.Errorf("first error")
	err2 := errors.NewWithCause(err1, "second error")

	DefaultLogger.SetLevel(LevelInfo)
	Info("test error", err2)
	Warnf("test %v", "error", err2)

	// Output:
	// 1982-03-15T12:56:14 [ INFO] test error ( logging_test.go:141 )
	//   > error: second error ( logging_test.go:137 )
	//   > cause by: first error
	// 1982-03-15T12:56:14 [ WARN] test error ( logging_test.go:142 )
	//   > error: second error ( logging_test.go:137 )
	//   > cause by: first error
}
