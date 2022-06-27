package depinject

// initializableReflectType is the type of Initializable interface.
//var initializableReflectType = reflect.TypeOf((*Initializable)(nil)).Elem()

// Initializable interface can be call when Context finish dependencies injection.
type Initializable interface {
	// AfterInject is calling when Context is started and dependencies injection.
	AfterInject() error
}
