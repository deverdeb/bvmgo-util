package depinject

import (
	"reflect"
	"testing"
)

type testStructure struct {
	attribute int
}

func (test *testStructure) method(_ string) int {
	return 0
}

var testIntrop1Type = reflect.TypeOf(testStructure{})
var ptrTestIntrop1Type = reflect.TypeOf((*testStructure)(nil))
var intType = reflect.TypeOf(1)
var intValue = 1
var ptrIntType = reflect.TypeOf(&intValue)

func Test_findStructType(t *testing.T) {

	type args struct {
		elemType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{name: "nil type", args: args{nil}, want: nil},
		{name: "testIntrop1Type type", args: args{testIntrop1Type}, want: testIntrop1Type},
		{name: "*testIntrop1Type type", args: args{ptrTestIntrop1Type}, want: testIntrop1Type},
		{name: "int type", args: args{intType}, want: nil},
		{name: "*int type", args: args{ptrIntType}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStructType(tt.args.elemType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStructType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNoPointerType(t *testing.T) {

	type args struct {
		elemType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{name: "nil type", args: args{nil}, want: nil},
		{name: "testIntrop1Type type", args: args{testIntrop1Type}, want: testIntrop1Type},
		{name: "*testIntrop1Type type", args: args{ptrTestIntrop1Type}, want: testIntrop1Type},
		{name: "int type", args: args{intType}, want: intType},
		{name: "*int type", args: args{ptrIntType}, want: intType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNoPointerType(tt.args.elemType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNoPointerType() = %v, want %v", got, tt.want)
			}
		})
	}
}
