package introsp

import (
	"github.com/deverdeb/bvmgo-util/errors"
	"reflect"
)

// Set method assigns the pointer to the value.
func Set(pointer interface{}, value interface{}) error {
	ptrValue := reflect.ValueOf(pointer)
	if ptrValue.Kind() != reflect.Ptr {
		return errors.New("failed to set value, require a pointer variable")
	}
	return SetReflectValue(ptrValue.Elem(), value)
}

// SetAttribute method assigns the structure attribute to the value.
func SetAttribute(structPointer interface{}, attribute string, value interface{}) error {
	ptrValue := reflect.ValueOf(structPointer)
	if ptrValue.Kind() != reflect.Ptr {
		return errors.New("failed to set attribute '%s', require a pointer of structure, unsupported type %s",
			attribute, TypeName(ptrValue.Type()))
	}
	structValue := ptrValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return errors.New("failed to set attribute '%s', require a pointer of structure, unsupported type %s",
			attribute, TypeName(ptrValue.Type()))
	}
	attributeValue := structValue.FieldByName(attribute)
	if !attributeValue.IsValid() {
		return errors.New("failed to set attribute '%s', attribute is not found on structure %s",
			attribute, TypeName(structValue.Type()))
	}
	if err := SetReflectValue(attributeValue, value); err != nil {
		return errors.NewWithCause(err, "failed to set attribute '%s', error found", attribute)
	}
	return nil
}

// SetReflectValue method assigns the reflect.Value to the value.
func SetReflectValue(elementValue reflect.Value, value interface{}) error {
	if value == nil {
		if !elementValue.CanSet() {
			return errors.New("failed to set value, variable can not be set (read only or not visible)")
		}
		elementValue.Set(reflect.Zero(elementValue.Type()))
	} else {
		// Adapt element and value types to match
		eltType, valueType, match := findMatchType(elementValue, reflect.ValueOf(value))
		if !match {
			return errors.New("failed to set value, value type '%s' can not assign to variable type '%s'",
				TypeName(valueType.Type()), TypeName(eltType.Type()))
		}
		if !eltType.CanSet() {
			return errors.New("failed to set value, variable can not be set (read only or not visible)")
		}
		eltType.Set(valueType)
	}
	return nil
}

// findMatchType find types for element and value.
// Types much verified valType.Type().AssignableTo(eltType.Type()).
func findMatchType(elementType reflect.Value, valueType reflect.Value) (eltType reflect.Value, valType reflect.Value, match bool) {
	eltType = elementType
	valType = valueType
	match = false
	// Value type is assignable to element type : return types
	if valType.Type().AssignableTo(eltType.Type()) {
		match = true
	} else if eltType.Kind() != reflect.Ptr && valType.Kind() == reflect.Ptr &&
		valType.Elem().Type().AssignableTo(eltType.Type()) {
		// Remove value pointer
		valType = valType.Elem()
		match = true
	}
	return
}
