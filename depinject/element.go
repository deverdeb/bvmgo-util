package depinject

import (
	"fmt"
	"reflect"
)

// elementStatus is the state of element in context
type elementStatus int

const (
	// Uninitialized is state of uninitialized elements
	Uninitialized elementStatus = iota
	// InInitialization is state of elements during initialization process
	InInitialization
	// Initialized is state of initialized elements
	Initialized
)

func (status elementStatus) ToString() string {
	switch status {
	case Uninitialized:
		return "Uninitialized"
	case InInitialization:
		return "InInitialization"
	case Initialized:
		return "Initialized"
	default:
		return "!!!Unknown status!!!"
	}
}

// elementInformation contains information about context element
type elementInformation struct {
	// eltType is the element type
	eltType reflect.Type
	// name is the element name
	name string
	// elementStatus is the element status
	status elementStatus
	// value is the element value
	value interface{}
}

func (element *elementInformation) ToString() string {
	return fmt.Sprintf("[type=%s, name='%s', status=%s]", element.eltType.Name(), element.name, element.status.ToString())
}
