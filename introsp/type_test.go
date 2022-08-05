package introsp

import (
	"reflect"
	"testing"
)

type TestStructureForTypeNameTest struct {
}

func (t TestStructureForTypeNameTest) Method1(string) int {
	return 0
}
func (t *TestStructureForTypeNameTest) Method2(int) string {
	return ""
}

func Func1ForTypeNameTest(string, int, float32) {
}
func Func2ForTypeNameTest(TestStructureForTypeNameTest) any {
	return nil
}
func Func3ForTypeNameTest() (string, error) {
	return "", nil
}

func TestValueTypeToString(t *testing.T) {
	type _ struct {
		elemType reflect.Type
	}
	tests := []struct {
		name string
		arg  any
		want string
	}{
		{name: "int type", arg: 123, want: "int"},
		{name: "bool type", arg: true, want: "bool"},
		{name: "string type", arg: "azerty", want: "string"},
		{name: "slice type", arg: make([]string, 0, 0), want: "[]string"},
		{name: "string type", arg: TestStructureForTypeNameTest{}, want: "TestStructureForTypeNameTest"},
		{name: "string type", arg: nil, want: "<nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueTypeToString(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeToString(t *testing.T) {
	channel := make(chan string)
	testStruct := TestStructureForTypeNameTest{}
	array := [3]int{1, 2, 3}
	slice := make([]int, 0, 3)
	mapStrInt := make(map[string]int)
	method1, _ := reflect.TypeOf(testStruct).MethodByName("Method1")
	method2, _ := reflect.TypeOf(&testStruct).MethodByName("Method2")

	type _ struct {
		elemType reflect.Type
	}
	tests := []struct {
		name string
		arg  reflect.Type
		want string
	}{
		{name: "int type", arg: reflect.TypeOf(123), want: "int"},
		{name: "bool type", arg: reflect.TypeOf(true), want: "bool"},
		{name: "string type", arg: reflect.TypeOf("azerty"), want: "string"},
		{name: "array type", arg: reflect.TypeOf(array), want: "[3]int"},
		{name: "slice type", arg: reflect.TypeOf(slice), want: "[]int"},
		{name: "map type", arg: reflect.TypeOf(mapStrInt), want: "map[string]int"},
		{name: "chan type", arg: reflect.TypeOf(channel), want: "chan string"},
		{name: "struct type", arg: reflect.TypeOf(testStruct), want: "TestStructureForTypeNameTest"},
		{name: "func type 1", arg: reflect.TypeOf(Func1ForTypeNameTest), want: "func (string, int, float32)"},
		{name: "func type 2", arg: reflect.TypeOf(Func2ForTypeNameTest), want: "func (TestStructureForTypeNameTest) any"},
		{name: "func type 3", arg: reflect.TypeOf(Func3ForTypeNameTest), want: "func () (string, error)"},
		{name: "method type 1", arg: method1.Type, want: "func (TestStructureForTypeNameTest, string) int"},
		{name: "method type 2", arg: method2.Type, want: "func (*TestStructureForTypeNameTest, int) string"},
		{name: "method type 3", arg: reflect.TypeOf(TestStructureForTypeNameTest.Method1), want: "func (TestStructureForTypeNameTest, string) int"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeToString(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMethodToString(t *testing.T) {
	testStruct := TestStructureForTypeNameTest{}
	method1, _ := reflect.TypeOf(testStruct).MethodByName("Method1")
	method2, _ := reflect.TypeOf(&testStruct).MethodByName("Method2")

	type _ struct {
		elemType reflect.Type
	}
	tests := []struct {
		name string
		arg  reflect.Method
		want string
	}{
		{name: "structure method type", arg: method1, want: "func (TestStructureForTypeNameTest) Method1(string) int"},
		{name: "pointer method type", arg: method2, want: "func (*TestStructureForTypeNameTest) Method2(int) string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MethodToString(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
