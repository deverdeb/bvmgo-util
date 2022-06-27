package logs

import (
	"fmt"
	"github.com/deverdeb/bvmgo-util/errors"
	"strings"
	"testing"
)

func TestFormatError(t *testing.T) {
	err := fmt.Errorf("big error for test !!!!")
	result := FormatError(err, 5)
	if result != "big error for test !!!!" {
		t.Errorf("FormatError() = %v, want %v", result, "big error for test !!!!")
	}
}

func TestFormatError_WithCause(t *testing.T) {
	err1 := fmt.Errorf("initial test error")
	err2 := errors.NewWithCause(err1, "second test error")
	err3 := errors.NewWithCause(err2, "third test error")
	result1 := FormatError(err3, 5)
	fmt.Println(result1)
	if !strings.Contains(result1, "initial test error") {
		t.Errorf("FormatError() = '%v', want contains '%v'", result1, "initial test error")
	}
	result2 := FormatError(err3, 2)
	if strings.Contains(result2, "initial test error") {
		t.Errorf("FormatError() = '%v', want does not contain '%v'", result2, "initial test error")
	}
	if !strings.Contains(result1, "second test error") {
		t.Errorf("FormatError() = '%v', want contains '%v'", result1, "second test error")
	}
	result3 := FormatError(err3, 1)
	if strings.Contains(result3, "second test error") {
		t.Errorf("FormatError() = '%v', want does not contain '%v'", result3, "second test error")
	}
	if !strings.Contains(result3, "third test error") {
		t.Errorf("FormatError() = '%v', want contains '%v'", result3, "third test error")
	}
	result4 := FormatError(err3, 0)
	if result4 != "..." {
		t.Errorf("FormatError() = '%v', want '%v'", result4, "...")
	}
}
