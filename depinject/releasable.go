package depinject

import "reflect"

// releasableReflectType is the type of Releasable interface.
var releasableReflectType = reflect.TypeOf((*Releasable)(nil)).Elem()

// Releasable interface can be released when Context is stopped.
type Releasable interface {
	// Release is calling when Context is stopped.
	Release()
}
