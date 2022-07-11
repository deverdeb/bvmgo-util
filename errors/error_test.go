package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrorNew(t *testing.T) {
	type args struct {
		format     string
		attributes []interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantMessage string
	}{
		{
			name: "build an Error",
			args: args{
				format:     "my %s error",
				attributes: []interface{}{"test"},
			},
			wantMessage: "my test error ( at github.com/deverdeb/bvmgo-util/errors.TestErrorNew.func1:29 )",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.format, tt.args.attributes...); err.Error() != tt.wantMessage {
				t.Errorf("Error() error = '%v', wantErr '%v'", err, tt.wantMessage)
			}
		})
	}
}

func TestErrorNewWithCause(t *testing.T) {
	type args struct {
		cause      error
		format     string
		attributes []interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantMessage string
	}{
		{
			name: "build an Error with cause",
			args: args{
				cause:      fmt.Errorf("cause error"),
				format:     "my %s error",
				attributes: []interface{}{"test"},
			},
			wantMessage: "my test error ( at github.com/deverdeb/bvmgo-util/errors.TestErrorNewWithCause.func1:60 )\n" +
				"    > cause by: cause error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewWithCause(tt.args.cause, tt.args.format, tt.args.attributes...); err.Error() != tt.wantMessage {
				t.Errorf("ErrorWithCause() error = '%v', wantErr '%v'", err, tt.wantMessage)
			}
		})
	}
}

func Test_customError_Error(t *testing.T) {
	type fields struct {
		message string
		cause   error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Error to string",
			fields: fields{
				message: "my test error",
				cause:   nil,
			},
			want: "my test error",
		}, {
			name: "Error with cause to string",
			fields: fields{
				message: "my test error",
				cause:   fmt.Errorf("cause error"),
			},
			want: "my test error\n    > cause by: cause error",
		}, {
			name: "Error without file information",
			fields: fields{
				message: "my test error",
				cause:   nil,
			},
			want: "my test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &customError{
				message: tt.fields.message,
				cause:   tt.fields.cause,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestErrorWrap(t *testing.T) {
	type args struct {
		cause      error
		format     string
		attributes []interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantMessage string
	}{
		{
			name: "wrap an Error",
			args: args{
				cause: fmt.Errorf("error message"),
			},
			wantMessage: "error message ( at github.com/deverdeb/bvmgo-util/errors.TestErrorWrap.func1:144 )\n" +
				"    > cause by: error message",
		},
		{
			name: "wrap a traceable Error",
			args: args{
				cause: New("error message"),
			},
			wantMessage: "error message ( at github.com/deverdeb/bvmgo-util/errors.TestErrorWrap.func1:144 )\n" +
				"    > cause by: error message ( at github.com/deverdeb/bvmgo-util/errors.TestErrorWrap:134 )",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.cause); err.Error() != tt.wantMessage {
				t.Errorf("Wrap() error = '%v', wantErr '%v'", err, tt.wantMessage)
			}
		})
	}
}

func TestErrorUnwrap(t *testing.T) {
	type Wrapper interface {
		Unwrap() error
	}
	err1 := New("error message")
	err2, _ := Wrap(err1).(Wrapper)
	err3 := err2.Unwrap()
	if err1 != err3 {
		t.Errorf("Unwrap() error = '%v', wantErr '%v'", err3, err1)
	}
}

func TestError_Cause(t *testing.T) {
	err1 := New("error message")
	err2, _ := Wrap(err1).(TraceableError)
	if err1 != err2.Cause() {
		t.Errorf("Cause() error = '%v', wantErr '%v'", err2.Cause(), err1)
	}
}

func TestError_File(t *testing.T) {
	err := New("error message").(TraceableError)
	if !strings.Contains(err.File(), "bvmgo-util/errors/error_test.go") {
		t.Errorf("File() = '%v', want contains '%v'", err.File(), "bvmgo-util/errors/error_test.go")
	}
}

func TestError_Function(t *testing.T) {
	err := New("error message").(TraceableError)
	if err.Function() != "github.com/deverdeb/bvmgo-util/errors.TestError_Function" {
		t.Errorf("Function() = '%v', want contains '%v'", err.Function(), "github.com/deverdeb/bvmgo-util/errors.TestError_Function")
	}
}

func TestError_Line(t *testing.T) {
	err := New("error message").(TraceableError)
	if err.Line() != 185 {
		t.Errorf("Line() = '%v', want '%v'", err.Line(), 185)
	}
}
