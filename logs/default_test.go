package logs

import (
	"fmt"
	"github.com/deverdeb/bvmgo-util/errors"
	"io"
	"log"
	"os"
	"testing"
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
	// 1982-03-15T12:56:14 [TRACE] hello everybody ( default_test.go:40 )
	// 1982-03-15T12:56:14 [TRACE] hello 123456 ( default_test.go:41 )
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
	// 1982-03-15T12:56:14 [DEBUG] hello everybody ( default_test.go:57 )
	// 1982-03-15T12:56:14 [DEBUG] hello world ( default_test.go:59 )
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
	// 1982-03-15T12:56:14 [ INFO] hello everybody ( default_test.go:75 )
	// 1982-03-15T12:56:14 [ INFO] hello false ( default_test.go:77 )
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
	// 1982-03-15T12:56:14 [ WARN] hello everybody ( default_test.go:93 )
	// 1982-03-15T12:56:14 [ WARN] hello world ( default_test.go:95 )
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
	// 1982-03-15T12:56:14 [ERROR] hello everybody ( default_test.go:111 )
	// 1982-03-15T12:56:14 [ERROR] hello world ( default_test.go:113 )
}

func ExampleFatal() {
	mockBeforeTest()
	defer restoreAfterTest()

	DefaultLogger().SetLevel(LevelError)
	Fatal("hello everybody")
	DefaultLogger().SetLevel(LevelFatal)
	Fatalf("hello %v", "world")

	// Output:
	// 1982-03-15T12:56:14 [FATAL] hello everybody ( default_test.go:129 )
	// exit with code 1
	// 1982-03-15T12:56:14 [FATAL] hello world ( default_test.go:131 )
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
	// 1982-03-15T12:56:14 [ INFO] test error ( default_test.go:148 )
	//   > error: second error ( default_test.go:144 )
	//   > cause by: first error
	// 1982-03-15T12:56:14 [ WARN] test error ( default_test.go:149 )
	//   > error: second error ( default_test.go:144 )
	//   > cause by: first error
}

func BenchmarkInfof(b *testing.B) {
	DefaultLogger().SetLevel(LevelInfo)
	for i := 0; i < b.N; i++ {
		Infof("test message %d %s %v", 12, "str", true)
	}
}

func BenchmarkInfof_WithError(b *testing.B) {
	DefaultLogger().SetLevel(LevelInfo)
	err1 := fmt.Errorf("first error")
	err2 := errors.NewWithCause(err1, "second error")
	for i := 0; i < b.N; i++ {
		Infof("test message %d %s %v", 12, "str", true, err2)
	}
}
