package introsp

import (
	"github.com/deverdeb/bvmgo-util/errors"
	"reflect"
)

// GetAttribute returns the value of structure attribute.
func GetAttribute(element interface{}, attribute string) (interface{}, error) {
	elementValue := reflect.ValueOf(element)
	attributeValue, find := findAttributeValue(elementValue, attribute)
	if !find {
		return nil, errors.New("failed to find attribute '%s' of '%s' type", attribute, TypeToString(elementValue.Type()))
	}
	if !attributeValue.CanInterface() {
		return nil, errors.New("failed to get attribute '%s' of '%s' type: unsupported type '%s' for result",
			attribute, TypeToString(elementValue.Type()), TypeToString(attributeValue.Type()))
	}
	return attributeValue.Interface(), nil
}

// findAttributeValue search and returns attribute value for a type value.
// Method tries to find value from fields (only for structure) and from methods.
// If type is a pointer, method tries to find from pointer and from pointed structure.
// findAttributeValue returns invalid value and false boolean if value is not found.
func findAttributeValue(typeValue reflect.Value, attributeName string) (reflect.Value, bool) {
	// search on type fields
	value, find := findAttributeValueFromField(typeValue, attributeName)
	if !find {
		// field not found, search on type methods
		value, find = findAttributeValueFromMethod(typeValue, attributeName)
		if !find && typeValue.Kind() == reflect.Ptr {
			// attribute is not found and type is pointer, try with pointed type
			if typeValue.Kind() == reflect.Ptr {
				value, find = findAttributeValue(typeValue.Elem(), attributeName)
			}
		}
	}
	return value, find
}

// findAttributeValueFromField search and returns attribute value for a type value.
// Method tries to find value from structure fields.
// findAttributeValueFromField returns invalid value and false boolean if field is not found or if entry value type is not a structure.
func findAttributeValueFromField(typeValue reflect.Value, attributeName string) (reflect.Value, bool) {
	// test entry value type (method supports only structure types)
	if typeValue.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	// search value from field
	fieldValue := typeValue.FieldByName(attributeName)
	if !fieldValue.IsValid() {
		return reflect.Value{}, false
	}
	return fieldValue, true
}

// findAttributeValueFromMethod search and returns attribute value for a type value.
// Method tries to find value from methods.
// findAttributeValue checks methods inputs and outputs (no parameters and only one returned value).
// findAttributeValue returns invalid value and false boolean if method is not found.
func findAttributeValueFromMethod(typeValue reflect.Value, attributeName string) (reflect.Value, bool) {
	// search value from method
	methodValue := typeValue.MethodByName(attributeName)
	if !methodValue.IsValid() {
		return reflect.Value{}, false
	}
	// method require no parameter
	methodType := methodValue.Type()
	if methodType.NumIn() != 0 {
		return reflect.Value{}, false
	}
	// method require only one result
	if methodType.NumOut() != 1 {
		return reflect.Value{}, false
	}
	// call method and return result
	methodReturnValues := methodValue.Call(nil)
	// verify result value
	if len(methodReturnValues) < 1 {
		return reflect.Value{}, false
	}
	value := methodReturnValues[0]
	if value.IsValid() {
		return value, true
	}
	return reflect.Value{}, false
}
