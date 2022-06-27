package depinject

import (
	"reflect"
)

// findStructType extract the "structure" type (Kind = Struct).
// If type is nil, return nil.
// If type is a structure, return the type.
// If type is a pointer, return the pointed type (if it is a structure type).
// Else return nil.
func findStructType(elemType reflect.Type) reflect.Type {
	if elemType == nil {
		return nil
	} else {
		switch elemType.Kind() {
		case reflect.Struct:
			return elemType
		case reflect.Ptr:
			return findStructType(elemType.Elem())
		default:
			return nil
		}
	}
}

// findNoPointerType extract the type.
// If type is nil, return nil.
// If type is a pointer, return the pointed type.
// Else return type.
func findNoPointerType(elemType reflect.Type) reflect.Type {
	if elemType == nil {
		return nil
	} else {
		switch elemType.Kind() {
		case reflect.Ptr:
			return findNoPointerType(elemType.Elem())
		default:
			return elemType
		}
	}
}
