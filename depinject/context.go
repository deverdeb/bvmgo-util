package depinject

import (
	"github.com/deverdeb/bvmgo-util/errors"
	"github.com/deverdeb/bvmgo-util/introsp"
	"reflect"
	"strings"
)

const InjectTag string = "inject"

// Context is a container for application structures.
// It injects dependencies by type or name.
type Context struct {
	// Is context started ?
	started bool
	// elements contains all context elements.
	elements []*elementInformation
	// initializedElements contains all initialized elements.
	// Elements are ordered by initialization order.
	initializedElements []*elementInformation
}

// CreateContext build an empty context instance.
func CreateContext() Context {
	return Context{
		started:             false,
		elements:            make([]*elementInformation, 0),
		initializedElements: make([]*elementInformation, 0),
	}
}

// GlobalContext is a default application global context.
var GlobalContext = CreateContext()

// Add an element to context.
func (context *Context) Add(element interface{}) error {
	if element == nil {
		return errors.New("context does not support nil element")
	}
	return context.AddWithName(element, introsp.TypeName(reflect.TypeOf(element)))
}

// AddWithName add an element with a name to context.
// Method returns error if another element exists with same name and type.
func (context *Context) AddWithName(element interface{}, name string) error {
	if element == nil {
		return errors.New("context does not support nil element")
	}
	// check if not exists another element with same name and type
	typeOfElement := reflect.TypeOf(element)
	alreadyExistElements := context.getElementsByNameAndType(name, typeOfElement)
	if len(alreadyExistElements) > 0 {
		return errors.New("cannot add '%s' element, another element exists with same name and type: %s",
			name, alreadyExistElements[0].ToString())
	}
	information := &elementInformation{
		eltType: typeOfElement,
		name:    name,
		status:  Uninitialized,
		value:   element,
	}
	context.elements = append(context.elements, information)
	if context.started {
		return context.initializeElement(information)
	}
	return nil
}

// Start inject dependencies and call `AfterInject()` methods of context structures.
func (context *Context) Start() error {
	// Inject dependencies
	for _, element := range context.elements {
		err := context.initializeElement(element)
		if err != nil {
			context.Stop()
			return errors.NewWithCause(err, "failed to start context, "+
				"error during '%s' element initialization", element.name)
		}
	}
	// Context is started
	context.started = true
	return nil
}

// Stop call `Release()` methods of context structures.
func (context *Context) Stop() {
	for _, element := range context.initializedElements {
		context.releaseElement(element)
	}
	context.initializedElements = make([]*elementInformation, 0)
	context.started = false
}

func (context *Context) initializeElement(information *elementInformation) error {
	if information.status == InInitialization {
		return errors.New("failed to initialized '%s' element, potential dependency loop", information.ToString())
	}
	if information.status == Uninitialized {
		// Verify dependencies
		err := context.injectDependencies(information)
		if err != nil {
			context.releaseElement(information)
			return errors.NewWithCause(err, "failed to inject dependencies of '%s' element", information.ToString())
		}
		// Execute process after injection
		err = context.callAfterInject(information)
		if err != nil {
			context.releaseElement(information)
			return errors.NewWithCause(err, "failed to initialized '%s' element after dependencies injection", information.ToString())
		}
	}
	return nil
}

// injectDependencies injects dependencies from context to element.
func (context *Context) injectDependencies(information *elementInformation) error {
	information.status = InInitialization
	eltStructType := findStructType(information.eltType)
	if eltStructType == nil {
		// element is not a structure: no injections
		information.status = Initialized
		return nil
	}
	// loop on element fields to find fields with injection tag ("inject")
	fieldsNumber := eltStructType.NumField()
	for fieldIndex := 0; fieldIndex < fieldsNumber; fieldIndex++ {
		field := eltStructType.Field(fieldIndex)
		fieldType := findNoPointerType(field.Type)
		tagValue, ok := field.Tag.Lookup(InjectTag)
		if ok {
			dependencyName := strings.TrimSpace(tagValue)
			var dependency *elementInformation
			var err error
			if dependencyName == "" {
				dependency, err = context.getElementByType(fieldType)
				if err != nil {
					return errors.NewWithCause(err, "failed to find '%s' dependency (by type: %s) of '%s' element",
						field.Name, introsp.TypeName(fieldType), information.ToString())
				} else if dependency == nil {
					return errors.New("missing '%s' dependency (by type: %s) of '%s' element",
						field.Name, introsp.TypeName(fieldType), information.ToString())
				}
			} else {
				dependency, err = context.getElementByName(dependencyName)
				if err != nil {
					return errors.NewWithCause(err, "failed to find '%s' dependency (by name: %s) of '%s' element",
						field.Name, dependencyName, information.ToString())
				} else if dependency == nil {
					return errors.New("missing '%s' dependency (by name: %s) of '%s' element",
						field.Name, dependencyName, information.ToString())
				}
			}
			if dependency.status != Initialized {
				err = context.initializeElement(dependency)
				if err != nil {
					return errors.NewWithCause(err, "failed to initialized '%s' dependency of '%s' element",
						field.Name, information.ToString())
				}
			}
			err = introsp.SetAttribute(information.value, field.Name, dependency.value)
			if err != nil {
				return errors.NewWithCause(err, "failed to initialized '%s' dependency of '%s' element, field cannot be set",
					field.Name, information.ToString())
			}
		}
	}
	context.initializedElements = append(context.initializedElements, information)
	information.status = Initialized
	return nil
}

// callAfterInject finalize element initialization by call Initializable.AfterInject() method if exists.
func (context *Context) callAfterInject(information *elementInformation) error {
	if information == nil || information.value == nil || information.status != Initialized {
		return nil
	}
	initializable, ok := information.value.(Initializable)
	if ok {
		return initializable.AfterInject()
	}
	return nil
}

// releaseElement release a contexte element.
// Call Releasable.Release() method if element value implements Releasable interface.
func (context *Context) releaseElement(information *elementInformation) {
	context.callReleaseMethod(information)
	context.removeDependencies(information)
}

// callReleaseMethod call Releasable.Release() method if element implements Releasable interface.
func (context *Context) callReleaseMethod(information *elementInformation) {
	if information == nil || information.value == nil || information.status != Initialized {
		return
	}
	releasable, ok := information.value.(Releasable)
	if ok {
		releasable.Release()
	}
}

// removeDependencies removes all element dependencies
func (context *Context) removeDependencies(information *elementInformation) {
	eltStructType := findStructType(information.eltType)
	if eltStructType != nil {
		fieldsNumber := eltStructType.NumField()
		for fieldIndex := 0; fieldIndex < fieldsNumber; fieldIndex++ {
			field := eltStructType.Field(fieldIndex)
			_, ok := field.Tag.Lookup(InjectTag)
			if ok {
				_ = introsp.SetAttribute(information.value, field.Name, nil)
			}
		}
	}
	// FIXME ...
	// ...> context.initializedElements = remove(context.initializedElements, information)
	information.status = Uninitialized
}

func (context *Context) getElementByType(eltType reflect.Type) (*elementInformation, error) {
	if eltType == nil {
		return nil, nil
	}
	finds := make([]*elementInformation, 0)
	for _, element := range context.elements {
		if element.eltType == eltType ||
			element.eltType.AssignableTo(eltType) ||
			(element.eltType.Elem() != nil && element.eltType.Elem().AssignableTo(eltType)) {
			finds = append(finds, element)
		}
	}
	if len(finds) == 1 {
		return finds[0], nil
	} else if len(finds) == 0 {
		return nil, nil
	} else {
		eltsInfo := ""
		for _, element := range finds {
			if len(eltsInfo) != 0 {
				eltsInfo = eltsInfo + ", "
			}
			eltsInfo += element.ToString()
		}
		return nil, errors.New("too many elements for type '%s': %s", introsp.TypeName(eltType), eltsInfo)
	}
}

// GetByType returns element with parameter type.
// Method returns error if element is not found or if context contains more than one element with the type.
func (context *Context) GetByType(eltType reflect.Type) (interface{}, error) {
	if eltType == nil {
		return nil, errors.New("cannot find element with nil type")
	}
	element, err := context.getElementByType(eltType)
	if err != nil {
		return nil, err
	} else if element == nil {
		return nil, errors.NewWithCause(err, "cannot find element with '%s' type", introsp.TypeName(eltType))
	} else {
		return context.extractElementValue(element)
	}
}

// getElementByName search element with the parameter name.
// Method returns error if more than one element is found.
// Method returns nil if no element is found.
func (context *Context) getElementByName(name string) (*elementInformation, error) {
	finds := context.getElementsByName(name)
	if len(finds) == 1 {
		return finds[0], nil
	} else if len(finds) == 0 {
		return nil, nil
	} else {
		eltsInfo := ""
		for _, element := range finds {
			if len(eltsInfo) != 0 {
				eltsInfo = eltsInfo + ", "
			}
			eltsInfo += element.ToString()
		}
		return nil, errors.New("too many elements for name '%s': %s", name, eltsInfo)
	}
}

// getElementsByName search all elements with the parameter name.
// Method empty slice if no element is found.
func (context *Context) getElementsByName(name string) []*elementInformation {
	finds := make([]*elementInformation, 0)
	for _, element := range context.elements {
		if element.name == name {
			finds = append(finds, element)
		}
	}
	return finds
}

// getElementsByNameAndType search all elements with the parameter name and the parameter type.
// Method empty slice if no element is found.
func (context *Context) getElementsByNameAndType(name string, eltType reflect.Type) []*elementInformation {
	finds := make([]*elementInformation, 0)
	findsByName := context.getElementsByName(name)
	for _, element := range findsByName {
		if element.eltType == eltType ||
			element.eltType.AssignableTo(eltType) ||
			(element.eltType.Elem() != nil && element.eltType.Elem().AssignableTo(eltType)) {
			finds = append(finds, element)
		}
	}
	return finds
}

// GetByName returns element with parameter name.
// Method returns error if element is not found or if context contains more than one element with the name.
// Note: Use GetByNameAndType method to get element by name and type.
func (context *Context) GetByName(name string) (interface{}, error) {
	element, err := context.getElementByName(name)
	if err != nil {
		return nil, err
	} else if element == nil {
		return nil, errors.NewWithCause(err, "cannot find element with '%s' name", name)
	} else {
		return context.extractElementValue(element)
	}
}

// GetByNameAndType returns element with parameters name and type.
// Method returns error if element is not found or if context contains more than one element with the name and the type.
func (context *Context) GetByNameAndType(name string, eltType reflect.Type) (interface{}, error) {
	elements := context.getElementsByNameAndType(name, eltType)
	nbElements := len(elements)
	if nbElements == 1 {
		return context.extractElementValue(elements[0])
	} else if len(elements) <= 0 {
		return nil, errors.New("cannot find element with '%s' name and '%s' type",
			name, introsp.TypeName(eltType))
	} else {
		eltsInfo := ""
		for _, element := range elements {
			if len(eltsInfo) != 0 {
				eltsInfo = eltsInfo + ", "
			}
			eltsInfo += element.ToString()
		}
		return nil, errors.New("too many elements for name '%s' and type '%s': %s",
			name, introsp.TypeName(eltType), eltsInfo)
	}
}

func (context *Context) extractElementValue(element *elementInformation) (interface{}, error) {
	if element.status == Uninitialized && context.started {
		err := context.initializeElement(element)
		if err != nil {
			return nil, errors.NewWithCause(err, "cannot return '%s' element, failed to initialized", element.name)
		}
	}
	return element.value, nil

}
