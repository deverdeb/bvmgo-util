package introsp

import (
	"reflect"
	"testing"
)

func TestTypeName(t *testing.T) {
	type testStructure struct {
	}
	channel := make(chan string)
	testStruct := testStructure{}
	//var ptr uintptr = 132
	//unsafePtr := unsafe.Pointer(ptr)
	array := [3]int{1, 2, 3}
	slice := make([]int, 0, 3)
	mapStrInt := make(map[string]int)

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
		{name: "struct type", arg: reflect.TypeOf(testStruct), want: "testStructure"},
		//{name: "unsafe pointer type", arg: reflect.TypeOf(unsafePtr), want: "unsafe.Pointer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeName(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
