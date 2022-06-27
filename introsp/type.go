package introsp

import (
	"fmt"
	"reflect"
)

// TypeName return the type name.
func TypeName(t reflect.Type) string {
	if t == nil {
		return "<nil>"
	}
	switch t.Kind() {
	case reflect.Invalid:
		return "<invalid>"
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), TypeName(t.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", TypeName(t.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", TypeName(t.Key()), TypeName(t.Elem()))
	case reflect.Chan:
		return "chan " + TypeName(t.Elem())
	case reflect.Func:
		return "func " + TypeName(t.Elem()) + "(...)"
	case reflect.Ptr:
		return "*" + TypeName(t.Elem())
	case reflect.UnsafePointer:
		return "unsafe." + t.Name()
	default:
		return t.Name()
	}
}
