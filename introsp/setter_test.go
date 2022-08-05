package introsp

import (
	"fmt"
	"strings"
	"testing"
)

type TestSetterInterface interface {
	method(testParam string) int
}
type TestSetterStruct struct {
	a int
	b string
	c bool
}

func (test TestSetterStruct) method(testParam string) int {
	return len(testParam) * 2
}

func TestSet_bool(t *testing.T) {
	var variable = false
	if err := Set(&variable, true); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if !variable {
		t.Errorf("Set() variable = %v, want = %v", variable, true)
	}
}
func TestSet_int(t *testing.T) {
	var variable = 0
	var value = 1234
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_int8(t *testing.T) {
	var variable int8 = 0
	var value int8 = 123
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_int64(t *testing.T) {
	var variable int64 = 0
	var value int64 = 123456
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_float32(t *testing.T) {
	var variable float32 = 0.
	var value float32 = 123.123
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_float64(t *testing.T) {
	var variable = 0.
	var value = 123.123456
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_string(t *testing.T) {
	var variable = ""
	var value = "test string"
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_struct(t *testing.T) {
	var variable = TestSetterStruct{a: 0, b: "first", c: false}
	var value = TestSetterStruct{a: 1234, b: "second", c: true}
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_pointerValue(t *testing.T) {
	var variable = TestSetterStruct{a: 0, b: "first", c: false}
	var value = TestSetterStruct{a: 1234, b: "second", c: true}
	// Call "Set" method with a pointer for value.
	if err := Set(&variable, &value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_pointerVariableAndPointerValue(t *testing.T) {
	var variable *TestSetterStruct = nil
	var value = &TestSetterStruct{a: 1234, b: "second", c: true}
	// Call "Set" method with variable pointer address.
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_interface(t *testing.T) {
	var variable TestSetterInterface = nil
	var value = &TestSetterStruct{a: 1234, b: "second", c: true}
	// Call "Set" method with variable pointer address.
	if err := Set(&variable, value); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != value {
		t.Errorf("Set() variable = %v, want = %v", variable, value)
	}
}
func TestSet_nilVariable(t *testing.T) {
	var value = &struct{ a int }{a: 1234}
	err := Set(nil, value)
	msg := "failed to set value, require a pointer variable"
	if err == nil || !strings.Contains(err.Error(), msg) {
		t.Errorf("Set() error = %v, want error with message \"%s\"", err, msg)
	}
}
func TestSet_nilValueToPointer(t *testing.T) {
	var variable TestSetterInterface = &TestSetterStruct{a: 1234, b: "second", c: true}
	if err := Set(&variable, nil); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != nil {
		t.Errorf("Set() variable = %v, want = %v", variable, nil)
	}
}
func TestSet_nilValueToStruct(t *testing.T) {
	variable := TestSetterStruct{a: 1234, b: "second", c: true}
	zeroValue := TestSetterStruct{} // Default value for type
	if err := Set(&variable, nil); err != nil {
		t.Errorf("Set() error = %v, want no Error", err)
	}
	if variable != zeroValue {
		t.Errorf("Set() variable = %v, want = %v", variable, zeroValue)
	}
}
func TestSet_badType(t *testing.T) {
	var variable = "azerty"
	var value = 1234
	err := Set(&variable, value)
	msg := "failed to set value, value type 'int' can not assign to variable type 'string'"
	if err == nil || !strings.Contains(err.Error(), msg) {
		t.Errorf("Set() error = %v, want error with message \"%s\"", err, msg)
	}
}

func TestSetAttribute_int(t *testing.T) {
	type TestSetStruct struct {
		AttribInt int
	}
	testObj := TestSetStruct{}
	if err := SetAttribute(&testObj, "AttribInt", 1234); err != nil {
		t.Errorf("SetAttribute() error = %v, want no Error", err)
	}
	if testObj.AttribInt != 1234 {
		t.Errorf("SetAttribute() variable = %v, want = %v", testObj.AttribInt, 1234)
	}
}
func TestSetAttribute_string(t *testing.T) {
	type TestSetStruct struct {
		AttribString string
	}
	testObj := TestSetStruct{}
	if err := SetAttribute(&testObj, "AttribString", "test test test"); err != nil {
		t.Errorf("SetAttribute() error = %v, want no Error", err)
	}
	if testObj.AttribString != "test test test" {
		t.Errorf("SetAttribute() variable = %v, want = %v", testObj.AttribString, "test test test")
	}
}
func TestSetAttribute_struct(t *testing.T) {
	type TestSetStruct struct {
		AttribStruct TestSetterStruct
	}
	testObj := TestSetStruct{}
	value := TestSetterStruct{a: 1234, b: "test", c: true}
	if err := SetAttribute(&testObj, "AttribStruct", value); err != nil {
		t.Errorf("SetAttribute() error = %v, want no Error", err)
	}
	if testObj.AttribStruct != value {
		t.Errorf("SetAttribute() variable = %v, want = %v", testObj.AttribStruct, value)
	}
}
func TestSetAttribute_pointer(t *testing.T) {
	type TestSetStruct struct {
		AttribStruct *TestSetterStruct
	}
	testObj := TestSetStruct{}
	value := TestSetterStruct{a: 1234, b: "test", c: true}
	if err := SetAttribute(&testObj, "AttribStruct", &value); err != nil {
		t.Errorf("SetAttribute() error = %v, want no Error", err)
	}
	if *testObj.AttribStruct != value {
		t.Errorf("SetAttribute() variable = %v, want = %v", testObj.AttribStruct, value)
	}
}
func TestSetAttribute_errors(t *testing.T) {
	type TestStruct struct {
		Field1 int
		field2 int
	}
	testObj := TestStruct{}
	intValue := 123

	type args struct {
		structure     interface{}
		attributeName string
		value         interface{}
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "assign value to structure field",
			args: args{
				structure:     &testObj,
				attributeName: "Field1",
				value:         123,
			},
			want: nil,
		},
		{
			name: "invalid - not a pointer",
			args: args{
				structure:     testObj,
				attributeName: "Field1",
				value:         546,
			},
			want: fmt.Errorf("failed to set attribute 'Field1', require a pointer of structure, unsupported type TestStruct"),
		},
		{
			name: "invalid - not a pointer of structure",
			args: args{
				structure:     &intValue,
				attributeName: "Field1",
				value:         546,
			},
			want: fmt.Errorf("failed to set attribute 'Field1', require a pointer of structure, unsupported type *int"),
		},
		{
			name: "invalid - bad type",
			args: args{
				structure:     &testObj,
				attributeName: "Field1",
				value:         true,
			},
			want: fmt.Errorf("failed to set value, value type 'bool' can not assign to variable type 'int'"),
		},
		{
			name: "invalid - private field",
			args: args{
				structure:     &testObj,
				attributeName: "field2",
				value:         356,
			},
			want: fmt.Errorf("failed to set attribute 'field2', error found"),
		},
		{
			name: "invalid - field is not found",
			args: args{
				structure:     &testObj,
				attributeName: "Field3",
				value:         789,
			},
			want: fmt.Errorf("failed to set attribute 'Field3', attribute is not found on structure TestStruct"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetAttribute(tt.args.structure, tt.args.attributeName, tt.args.value)
			if got != nil && tt.want != nil && !strings.Contains(got.Error(), tt.want.Error()) {
				t.Errorf("setStructureFieldWithValue() = \"%v\", want error contains \"%v\"", got, tt.want)
			} else if (got != nil || tt.want != nil) && (got == nil || tt.want == nil) {
				t.Errorf("setStructureFieldWithValue() = \"%v\", want error \"%v\"", got, tt.want)
			}
		})
	}
}
