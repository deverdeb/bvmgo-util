package errors

import (
	"fmt"
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
			wantMessage: "my test error ( at github.com/deverdeb/bvmgo-util/errors.TestErrorNew.func1:28 )",
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
			wantMessage: "my test error ( at github.com/deverdeb/bvmgo-util/errors.TestErrorNewWithCause.func1:59 )\n" +
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
			wantMessage: "error message ( at github.com/deverdeb/bvmgo-util/errors.TestErrorWrap.func1:143 )\n" +
				"    > cause by: error message",
		},
		{
			name: "wrap a traceable Error",
			args: args{
				cause: New("error message"),
			},
			wantMessage: "error message ( at github.com/deverdeb/bvmgo-util/errors.TestErrorWrap.func1:143 )\n" +
				"    > cause by: error message ( at github.com/deverdeb/bvmgo-util/errors.TestErrorWrap:133 )",
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
