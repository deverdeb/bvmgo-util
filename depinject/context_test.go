package depinject

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type interfaceContextTest interface {
	Method()
}

type structContextTest struct {
	init    bool
	release bool
}

func (test *structContextTest) AfterInject() error {
	test.init = true
	return nil
}
func (test *structContextTest) Release() {
	test.release = true
}
func (test *structContextTest) Method() {
	fmt.Println("test message")
}

func TestContext_Add(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	err := testContext.Add(obj1)
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	result, err := testContext.GetByType(reflect.TypeOf(obj1))
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Object = %v, want = %v", result, obj1)
	}
	if obj1.init {
		t.Errorf("obj1.init = %v, want = %v", obj1.init, false)
	}
	if obj1.release {
		t.Errorf("obj1.release = %v, want = %v", obj1.release, false)
	}
}
func TestContext_Add_ContextStarted(t *testing.T) {
	testContext := CreateContext()
	err := testContext.Start()
	if err != nil {
		t.Errorf("cannot start context, error found: %v", err)
	}
	defer testContext.Stop()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	err = testContext.Add(obj1)
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	result, err := testContext.GetByType(reflect.TypeOf(obj1))
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Object = %v, want = %v", result, obj1)
	}
	if !obj1.init {
		t.Errorf("obj1.init = %v, want = %v", obj1.init, true)
	}
	if obj1.release {
		t.Errorf("obj1.release = %v, want = %v", obj1.release, false)
	}
}

func TestContext_AddWithName_struct(t *testing.T) {
	testContext := CreateContext()
	obj1 := structContextTest{
		init:    false,
		release: false,
	}
	if err := testContext.AddWithName(obj1, "Obj1"); err != nil {
		t.Errorf("Failed to add structure element = %v, want no error", err)
	}
	if err := testContext.AddWithName(&obj1, "PtrObj1"); err != nil {
		t.Errorf("Failed to add pointer element = %v, want no error", err)
	}
}

func TestContext_AddWithName_values(t *testing.T) {
	testContext := CreateContext()
	intValue := 123
	floatValue := 123.123
	stringValue := "qwerty"
	if err := testContext.AddWithName(intValue, "intValue"); err != nil {
		t.Errorf("Failed to add int value = %v, want no error", err)
	}
	if err := testContext.AddWithName(floatValue, "floatValue"); err != nil {
		t.Errorf("Failed to add float value = %v, want no error", err)
	}
	if err := testContext.AddWithName(stringValue, "stringValue"); err != nil {
		t.Errorf("Failed to add string value = %v, want no error", err)
	}
}

func TestContext_AddWithName_nil(t *testing.T) {
	testContext := CreateContext()
	if err := testContext.AddWithName(nil, "nilValue"); err == nil {
		t.Errorf("Error() = %v, want nil element error", err)
	}
}

func TestContext_AddWithName_AlreadyExists(t *testing.T) {
	testContext := CreateContext()
	obj1 := structContextTest{
		init:    false,
		release: false,
	}
	if err := testContext.AddWithName(obj1, "myObject"); err != nil {
		t.Errorf("Failed to add structure = %v, want no error", err)
	}
	obj2 := structContextTest{
		init:    false,
		release: false,
	}
	err := testContext.AddWithName(obj2, "myObject")
	if err == nil {
		t.Errorf("Error() = %v, want already exists element error", err)
		return
	}
	if !strings.Contains(err.Error(), "another element exists with same name and type") {
		t.Errorf("Error() = \"%v\", want \"another element exists with same name and type\" error message", err)
	}
}

func TestContext_AddWithName_ContextStartedAndAlreadyExists(t *testing.T) {
	testContext := CreateContext()
	_ = testContext.Start()
	defer testContext.Stop()
	obj1 := structContextTest{
		init:    false,
		release: false,
	}
	if err := testContext.AddWithName(obj1, "myObject"); err != nil {
		t.Errorf("Failed to add structure = %v, want no error", err)
	}
	obj2 := structContextTest{
		init:    false,
		release: false,
	}
	err := testContext.AddWithName(obj2, "myObject")
	if err == nil {
		t.Errorf("Error() = %v, want already exists element error", err)
		return
	}
	if !strings.Contains(err.Error(), "another element exists with same name and type") {
		t.Errorf("Error() = \"%v\", want \"another element exists with same name and type\" error message", err)
	}
}

func TestContext_GetByName_struct(t *testing.T) {
	// create and launch context
	testContext := CreateContext()
	_ = testContext.Start()
	defer testContext.Stop()
	// add elements to context
	obj1 := structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.AddWithName(obj1, "Obj1")
	ptrObj1 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.AddWithName(ptrObj1, "PtrObj1")
	// test read elements
	// - struct
	result, err := testContext.GetByName("Obj1")
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Result = %v, want struct = %v", result, obj1)
	}
	// - pointer
	result, err = testContext.GetByName("PtrObj1")
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != ptrObj1 {
		t.Errorf("Result = %v, want pointer = %v", result, ptrObj1)
	}
}

func TestContext_GetByName_values(t *testing.T) {
	// create and launch context
	testContext := CreateContext()
	_ = testContext.Start()
	defer testContext.Stop()
	// add values to context
	intValue := 123
	_ = testContext.AddWithName(intValue, "intValue")
	floatValue := 123.123
	_ = testContext.AddWithName(floatValue, "floatValue")
	stringValue := "qwerty"
	_ = testContext.AddWithName(stringValue, "stringValue")
	// test read values
	// - int
	result, err := testContext.GetByName("intValue")
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != intValue {
		t.Errorf("Result = %v, want = %v", result, intValue)
	}
	// - float
	result, err = testContext.GetByName("floatValue")
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != floatValue {
		t.Errorf("Result = %v, want = %v", result, floatValue)
	}
	// - string
	result, err = testContext.GetByName("stringValue")
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != stringValue {
		t.Errorf("Result = %v, want = %v", result, stringValue)
	}
}

func TestContext_GetByName_NotFound(t *testing.T) {
	testContext := CreateContext()
	result, err := testContext.GetByName("obj2")
	if err == nil {
		t.Errorf("Error() = %v, want not found element error", err)
	}
	if result != nil {
		t.Errorf("Result = %v, want = %v", result, nil)
	}
}

func TestContext_GetByType(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.Add(obj1)
	result, err := testContext.GetByType(reflect.TypeOf(obj1))
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Result = %v, want = %v", result, obj1)
	}
	result, err = testContext.GetByType(reflect.TypeOf((*interfaceContextTest)(nil)).Elem())
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Result = %v, want = %v", result, obj1)
	}
}

func TestContext_GetByType_NotFound(t *testing.T) {
	type notExistElementStructTest struct {
	}
	testContext := CreateContext()
	obj2 := &notExistElementStructTest{}
	result, err := testContext.GetByType(reflect.TypeOf(obj2))
	if err == nil {
		t.Errorf("Error() = %v, want not found element error", err)
	}
	if result != nil {
		t.Errorf("Result = %v, want = %v", result, nil)
	}
}

func TestContext_GetByNameAndType(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.AddWithName(obj1, "MyObject")
	result, err := testContext.GetByNameAndType("MyObject", reflect.TypeOf(obj1))
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Result = %v, want = %v", result, obj1)
	}
	result, err = testContext.GetByNameAndType("MyObject", reflect.TypeOf((*interfaceContextTest)(nil)).Elem())
	if err != nil {
		t.Errorf("Error() = %v, want no error", err)
	}
	if result != obj1 {
		t.Errorf("Result = %v, want = %v", result, obj1)
	}
}

func TestContext_GetByNameAndType_NotFound(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.AddWithName(obj1, "MyObject")
	result, err := testContext.GetByNameAndType("TestObject", reflect.TypeOf(obj1))
	if err == nil {
		t.Errorf("Error() = %v, want not found element error", err)
	}
	if result != nil {
		t.Errorf("Result = %v, want = %v", result, nil)
	}
	type notExistElementStructTest struct {
	}
	obj2 := &notExistElementStructTest{}
	result, err = testContext.GetByNameAndType("MyObject", reflect.TypeOf(obj2))
	if err == nil {
		t.Errorf("Error() = %v, want not found element error", err)
	}
	if result != nil {
		t.Errorf("Result = %v, want = %v", result, nil)
	}
}

func TestContext_Start(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.AddWithName(obj1, "obj1")
	obj2 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.Add(obj2)
	type otherStructTest struct {
	}
	obj3 := &otherStructTest{}
	_ = testContext.Add(obj3)
	_ = testContext.AddWithName(123, "intValue")
	_ = testContext.AddWithName(123.123, "floatValue")
	_ = testContext.AddWithName("string", "stringValue")
	err := testContext.Start()
	if err != nil {
		t.Errorf("cannot start context, error found: %v", err)
	}
	defer testContext.Stop()
	// Verify elements are initialized
	if !obj1.init {
		t.Errorf("obj1.init = %v, want = %v", obj1.init, true)
	}
	if !obj2.init {
		t.Errorf("obj2.init = %v, want = %v", obj2.init, true)
	}
}

func TestContext_Stop(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.AddWithName(obj1, "obj1")
	obj2 := &structContextTest{
		init:    false,
		release: false,
	}
	_ = testContext.Add(obj2)
	type otherStructTest struct {
	}
	obj3 := &otherStructTest{}
	_ = testContext.Add(obj3)
	_ = testContext.AddWithName(123, "intValue")
	_ = testContext.AddWithName(123.123, "floatValue")
	_ = testContext.AddWithName("string", "stringValue")
	_ = testContext.Start()
	testContext.Stop()
	// Verify elements are released
	if !obj1.release {
		t.Errorf("obj1.release = %v, want = %v", obj1.release, true)
	}
	if !obj2.release {
		t.Errorf("obj2.release = %v, want = %v", obj2.release, true)
	}
}
