package introsp

import (
	"strings"
	"testing"
)

type TestGetterInterface interface {
	AttribInt() int
}

type TestGetterStruct1 struct {
	a int
}

func (test TestGetterStruct1) AttribInt() int {
	return test.a
}
func (test TestGetterStruct1) Method1(a int) {
	test.a = a * 5
}

type TestGetterStruct2 struct {
	a int
}

func (test *TestGetterStruct2) AttribInt() int {
	return test.a
}

func TestGetAttribute_int(t *testing.T) {
	type TestGetStruct struct {
		AttribInt int
	}
	testObj := TestGetStruct{AttribInt: 1234}
	value, err := GetAttribute(testObj, "AttribInt")
	if err != nil {
		t.Errorf("GetAttribute() error = %v, want no Error", err)
	}
	if value != 1234 {
		t.Errorf("GetAttribute() value = %v, want = %v", value, 1234)
	}
}

func TestGetAttribute_string(t *testing.T) {
	type TestGetStruct struct {
		AttribString string
	}
	testObj := TestGetStruct{AttribString: "test test test"}
	value, err := GetAttribute(testObj, "AttribString")
	if err != nil {
		t.Errorf("GetAttribute() error = %v, want no Error", err)
	}
	if value != "test test test" {
		t.Errorf("GetAttribute() value = %v, want = %v", value, "test test test")
	}
}

func TestGetAttribute_struct(t *testing.T) {
	type TestChildStruct struct {
		AttribString string
	}
	type TestGetStruct struct {
		AttribStruct TestChildStruct
	}
	testObj := TestGetStruct{AttribStruct: TestChildStruct{AttribString: "test test test"}}
	value, err := GetAttribute(testObj, "AttribStruct")
	if err != nil {
		t.Errorf("GetAttribute() error = %v, want no Error", err)
	}
	structValue, ok := value.(TestChildStruct)
	if !ok {
		t.Errorf("GetAttribute() value type = %v, want = %v", ValueTypeToString(structValue), ValueTypeToString(testObj.AttribStruct))
	}
	if structValue.AttribString != "test test test" {
		t.Errorf("GetAttribute() value = %v, want = %v", structValue.AttribString, "test test test")
	}
}

func TestGetAttribute_pointer(t *testing.T) {
	type TestGetStruct struct {
		AttribInt int
	}
	testObj := TestGetStruct{AttribInt: 1234}
	value, err := GetAttribute(&testObj, "AttribInt")
	if err != nil {
		t.Errorf("GetAttribute() error = %v, want no Error", err)
	}
	if value != 1234 {
		t.Errorf("GetAttribute() value = %v, want = %v", value, 1234)
	}
}

func TestGetAttribute_methodAccessor(t *testing.T) {
	var testObj TestGetterInterface = TestGetterStruct1{a: 1234}
	value, err := GetAttribute(testObj, "AttribInt")
	if err != nil {
		t.Errorf("GetAttribute() error = %v, want no Error", err)
	}
	if value != 1234 {
		t.Errorf("GetAttribute() value = %v, want = %v", value, 1234)
	}
}

func TestGetAttribute_methodAccessorFromPointer(t *testing.T) {
	var testObj TestGetterInterface = &TestGetterStruct2{a: 1234}
	value, err := GetAttribute(testObj, "AttribInt")
	if err != nil {
		t.Errorf("GetAttribute(...) error = %v, want no Error", err)
	}
	if value != 1234 {
		t.Errorf("GetAttribute(...) value = %v, want = %v", value, 1234)
	}
}

func TestGetAttribute_invalidMethodSignature(t *testing.T) {
	testObj := TestGetterStruct1{a: 1234}
	_, err := GetAttribute(&testObj, "Method1")
	msg := "failed to find attribute 'Method1' of '*TestGetterStruct1' type"
	if err == nil || !strings.Contains(err.Error(), msg) {
		t.Errorf("GetAttribute(...) error = %v, want error with message \"%s\"", err, msg)
	}
}
