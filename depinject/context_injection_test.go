package depinject

import (
	"fmt"
	"strings"
	"testing"
)

type interfaceInjectTest1 interface {
	Method()
}

type structInjectTest1 struct {
}

func (test *structInjectTest1) Method() {
	fmt.Println("test message")
}

type structInjectTest2 struct {
}

type structInjectTest3 struct {
	Obj1Struct1 interfaceInjectTest1 `inject:"obj1"`
	Obj2Struct1 *structInjectTest1   `inject:"obj2"`
	Obj3Struct2 *structInjectTest2   `inject:""`
}

func TestContext_Start_WithInjection(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structInjectTest1{}
	_ = testContext.AddWithName(obj1, "obj1")
	obj2 := &structInjectTest1{}
	_ = testContext.AddWithName(obj2, "obj2")
	obj3 := &structInjectTest2{}
	_ = testContext.Add(obj3)
	obj4 := &structInjectTest3{
		Obj1Struct1: nil,
		Obj2Struct1: nil,
		Obj3Struct2: nil,
	}
	_ = testContext.Add(obj4)
	err := testContext.Start()
	if err != nil {
		t.Errorf("cannot start context, error found: %v", err)
	}
	defer testContext.Stop()
	// Verify dependencies injections
	if obj4.Obj1Struct1 != obj1 {
		t.Errorf("obj3.Obj1Struct1 = %v, want = %v", obj4.Obj1Struct1, obj1)
	}
	if obj4.Obj2Struct1 != obj2 {
		t.Errorf("obj3.Obj2Struct1 = %v, want = %v", obj4.Obj2Struct1, obj2)
	}
	if obj4.Obj3Struct2 != obj3 {
		t.Errorf("obj3.Obj3Struct2 = %v, want = %v", obj4.Obj3Struct2, obj3)
	}
}

func TestContext_Start_WithInjectionOfNotInitializedElements(t *testing.T) {
	testContext := CreateContext()
	obj4 := &structInjectTest3{
		Obj1Struct1: nil,
		Obj2Struct1: nil,
		Obj3Struct2: nil,
	}
	_ = testContext.Add(obj4)
	obj1 := &structInjectTest1{}
	_ = testContext.AddWithName(obj1, "obj1")
	obj2 := &structInjectTest1{}
	_ = testContext.AddWithName(obj2, "obj2")
	obj3 := &structInjectTest2{}
	_ = testContext.Add(obj3)
	err := testContext.Start()
	if err != nil {
		t.Errorf("cannot start context, error found: %v", err)
	}
	defer testContext.Stop()
	// Verify dependencies injections
	if obj4.Obj1Struct1 != obj1 {
		t.Errorf("obj3.Obj1Struct1 = %v, want = %v", obj4.Obj1Struct1, obj1)
	}
	if obj4.Obj2Struct1 != obj2 {
		t.Errorf("obj3.Obj2Struct1 = %v, want = %v", obj4.Obj2Struct1, obj2)
	}
	if obj4.Obj3Struct2 != obj3 {
		t.Errorf("obj3.Obj3Struct2 = %v, want = %v", obj4.Obj3Struct2, obj3)
	}
}

func TestContext_Start_LoopInjection(t *testing.T) {
	testContext := CreateContext()
	var1 := struct {
		field2 interface{} `inject:"var2"`
	}{}
	_ = testContext.AddWithName(&var1, "var1")
	var2 := struct {
		field3 interface{} `inject:"var3"`
	}{}
	_ = testContext.AddWithName(&var2, "var2")
	var3 := struct {
		field1 interface{} `inject:"var1"`
	}{}
	_ = testContext.AddWithName(&var3, "var3")
	err := testContext.Start()
	if err == nil {
		t.Errorf("Error() = %v, want Error contains \"dependency loop\"", err)
	}
	if !strings.Contains(err.Error(), "dependency loop") {
		t.Errorf("Error() = %v, want contains \"dependency loop\"", err)
	}
}

func TestContext_Start_MissingDependency(t *testing.T) {
	testContext := CreateContext()
	var1 := struct {
		field2 interface{} `inject:"var2"`
	}{}
	_ = testContext.AddWithName(&var1, "var1")
	var2 := struct {
		field3 interface{} `inject:"var3"`
	}{}
	_ = testContext.AddWithName(&var2, "var2")
	err := testContext.Start()
	if err == nil {
		t.Errorf("Error() = %v, want Error contains \"missing 'field3' dependency\"", err)
	}
	if !strings.Contains(err.Error(), "missing 'field3' dependency") {
		t.Errorf("Error() = %v, want contains \"missing 'field3' dependency\"", err)
	}
}

func TestContext_Start_TooManyDependencies(t *testing.T) {
	testContext := CreateContext()
	type type2 struct{}
	var1 := struct {
		field *type2 `inject:""`
	}{}
	_ = testContext.AddWithName(&var1, "var1")
	_ = testContext.AddWithName(&type2{}, "var2-1")
	_ = testContext.AddWithName(&type2{}, "var2-2")
	err := testContext.Start()
	if err == nil {
		t.Errorf("Error() = %v, want Error contains \"missing 'field3' dependency\"", err)
	}
	if !strings.Contains(err.Error(), "too many elements for type 'type2'") {
		t.Errorf("Error() = %v, want contains \"too many elements for type 'type2'\"", err)
	}
}

func TestContext_Start_WithInjectionOfValues(t *testing.T) {
	testContext := CreateContext()
	var1 := struct {
		IntField    int     `inject:"intValue"`
		FloatField  float64 `inject:"floatValue"`
		StringField string  `inject:"stringValue"`
	}{}
	_ = testContext.Add(&var1)
	intValue := 123
	_ = testContext.AddWithName(intValue, "intValue")
	floatValue := 123.123
	_ = testContext.AddWithName(floatValue, "floatValue")
	stringValue := "qwerty"
	_ = testContext.AddWithName(stringValue, "stringValue")
	err := testContext.Start()
	if err != nil {
		t.Errorf("cannot start context, error found: %v", err)
	}
	defer testContext.Stop()
	// Verify dependencies injections
	if var1.IntField != intValue {
		t.Errorf("obj3.IntField = %v, want = %v", var1.IntField, intValue)
	}
	if var1.FloatField != floatValue {
		t.Errorf("obj3.FloatField = %v, want = %v", var1.FloatField, floatValue)
	}
	if var1.StringField != stringValue {
		t.Errorf("obj3.StringField = %v, want = %v", var1.StringField, stringValue)
	}
}

func TestContext_Stop_ReleaseInjections(t *testing.T) {
	testContext := CreateContext()
	obj1 := &structInjectTest1{}
	_ = testContext.AddWithName(obj1, "obj1")
	obj2 := &structInjectTest1{}
	_ = testContext.AddWithName(obj2, "obj2")
	obj3 := &structInjectTest2{}
	_ = testContext.Add(obj3)
	obj4 := &structInjectTest3{
		Obj1Struct1: nil,
		Obj2Struct1: nil,
		Obj3Struct2: nil,
	}
	_ = testContext.Add(obj4)
	_ = testContext.Start()
	testContext.Stop()
	// Verify elements are released
	if obj4.Obj1Struct1 != nil {
		t.Errorf("obj3.Obj1Struct1 = %v, want = %v", obj4.Obj1Struct1, nil)
	}
	if obj4.Obj2Struct1 != nil {
		t.Errorf("obj3.Obj2Struct1 = %v, want = %v", obj4.Obj2Struct1, nil)
	}
	if obj4.Obj3Struct2 != nil {
		t.Errorf("obj3.Obj3Struct2 = %v, want = %v", obj4.Obj3Struct2, nil)
	}
}
