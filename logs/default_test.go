package logs

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
	DefaultLogger().Output().SetOutput(os.Stdout)
	now = func() time.Time {
		return time.Date(1982, time.March, 15, 12, 56, 14, 123, time.Local)
	}
	exit = func(code int) {
		DefaultLogger().Output().Printf("exit with code %d", code)
	}
}

// restoreAfterTest restore the default "Now()" function and log writer.
func restoreAfterTest() {
	DefaultLogger().Output().SetOutput(previousWriter)
	now = time.Now
	exit = os.Exit
}

func ExampleTrace() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelTrace)
	Trace("hello everybody")
	Tracef("hello %d", 123456)
	DefaultLogger().SetLevel(LevelDebug)
	Trace("hidden log")
	DefaultLogger().SetLevel(LevelInfo)
	Tracef("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [TRACE] hello everybody ( default_test.go:39 )
	// 1982-03-15T12:56:14 [TRACE] hello 123456 ( default_test.go:40 )
}

func ExampleDebug() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelTrace)
	Debug("hello everybody")
	DefaultLogger().SetLevel(LevelDebug)
	Debugf("hello %s", "world")
	DefaultLogger().SetLevel(LevelInfo)
	Debug("hidden log")
	DefaultLogger().SetLevel(LevelWarn)
	Debugf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [DEBUG] hello everybody ( default_test.go:56 )
	// 1982-03-15T12:56:14 [DEBUG] hello world ( default_test.go:58 )
}

func ExampleInfo() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelTrace)
	Info("hello everybody")
	DefaultLogger().SetLevel(LevelInfo)
	Infof("hello %v", false)
	DefaultLogger().SetLevel(LevelWarn)
	Info("hidden log")
	DefaultLogger().SetLevel(LevelError)
	Infof("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ INFO] hello everybody ( default_test.go:74 )
	// 1982-03-15T12:56:14 [ INFO] hello false ( default_test.go:76 )
}

func ExampleWarn() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelInfo)
	Warn("hello everybody")
	DefaultLogger().SetLevel(LevelWarn)
	Warnf("hello %v", "world")
	DefaultLogger().SetLevel(LevelError)
	Warn("hidden log")
	DefaultLogger().SetLevel(LevelFatal)
	Warnf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ WARN] hello everybody ( default_test.go:92 )
	// 1982-03-15T12:56:14 [ WARN] hello world ( default_test.go:94 )
}

func ExampleError() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelInfo)
	Error("hello everybody")
	DefaultLogger().SetLevel(LevelError)
	Errorf("hello %v", "world")
	DefaultLogger().SetLevel(LevelFatal)
	Error("hidden log")
	DefaultLogger().SetLevel(LevelFatal)
	Errorf("hidden log format")

	// Output:
	// 1982-03-15T12:56:14 [ERROR] hello everybody ( default_test.go:110 )
	// 1982-03-15T12:56:14 [ERROR] hello world ( default_test.go:112 )
}

func ExampleFatal() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelError)
	Fatal("hello everybody")
	DefaultLogger().SetLevel(LevelFatal)
	Fatalf("hello %v", "world")

	// Output:
	// 1982-03-15T12:56:14 [FATAL] hello everybody ( default_test.go:128 )
	// exit with code 1
	// 1982-03-15T12:56:14 [FATAL] hello world ( default_test.go:130 )
	// exit with code 1
}

func ExampleLogger_Formatter() {
	mockBeforeTest()
	defer restoreAfterTest()

	err1 := fmt.Errorf("first error")
	err2 := errors.NewWithCause(err1, "second error")

	DefaultLogger().SetLevel(LevelInfo)
	Info("test error", err2)
	Warnf("test %v", "error", err2)

	// Output:
	// 1982-03-15T12:56:14 [ INFO] test error ( default_test.go:147 )
	//   > error: second error ( default_test.go:143 )
	//   > cause by: first error
	// 1982-03-15T12:56:14 [ WARN] test error ( default_test.go:148 )
	//   > error: second error ( default_test.go:143 )
	//   > cause by: first error
}
