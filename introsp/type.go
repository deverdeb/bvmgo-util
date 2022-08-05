package introsp

import (
	"fmt"
	"reflect"
)

// ValueTypeToString returns type name (string) of value.
func ValueTypeToString(value any) string {
	return TypeToString(reflect.TypeOf(value))
}

// TypeToString returns type name (string).
func TypeToString(t reflect.Type) string {
	if t == nil {
		return "<nil>"
	}
	switch t.Kind() {
	case reflect.Invalid:
		return "<invalid>"
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), TypeToString(t.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", TypeToString(t.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", TypeToString(t.Key()), TypeToString(t.Elem()))
	case reflect.Chan:
		return "chan " + TypeToString(t.Elem())
	case reflect.Func:
		return funcTypeToString(t)
	case reflect.Ptr:
		return "*" + TypeToString(t.Elem())
	case reflect.UnsafePointer:
		return "unsafe." + t.Name()
	case reflect.Interface:
		return interfaceTypeToString(t)
	default:
		return t.Name()
	}
}

// MethodToString return the method type name.
func MethodToString(method reflect.Method) string {
	methodType := method.Type
	name := "func (" + TypeToString(methodType.In(0)) + ") " + method.Name + "("
	for idx := 1; idx < methodType.NumIn(); idx++ {
		paramType := methodType.In(idx)
		if idx > 1 {
			name += ", "
		}
		name += TypeToString(paramType)
	}
	name += ")"
	if methodType.NumOut() == 1 {
		name += " " + TypeToString(methodType.Out(0))
	} else if methodType.NumOut() > 1 {
		name += " ("
		for idx := 0; idx < methodType.NumOut(); idx++ {
			paramType := methodType.Out(idx)
			if idx > 0 {
				name += ", "
			}
			name += TypeToString(paramType)
		}
		name += ")"
	}
	return name
}

func interfaceTypeToString(t reflect.Type) string {
	if t.Name() == "" {
		return "any"
	} else {
		return t.Name()
	}
}

func funcTypeToString(t reflect.Type) string {
	name := "func ("
	for idx := 0; idx < t.NumIn(); idx++ {
		paramType := t.In(idx)
		if idx > 0 {
			name += ", "
		}
		name += TypeToString(paramType)
	}
	name += ")"
	if t.NumOut() == 1 {
		name += " " + TypeToString(t.Out(0))
	} else if t.NumOut() > 1 {
		name += " ("
		for idx := 0; idx < t.NumOut(); idx++ {
			paramType := t.Out(idx)
			if idx > 0 {
				name += ", "
			}
			name += TypeToString(paramType)
		}
		name += ")"
	}
	return name
}
