package properties

import (
	"testing"
)

func TestReadFromFileToMap(t *testing.T) {
	properties, err := ReadFromFileToMap("test.properties")
	if err != nil {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", err, nil)
		return
	}
	if properties["key-without-value-1"] != "" {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", properties["key-without-value-1"], "")
	}
	if properties["key-without-value-2"] != "" {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", properties["key-without-value-2"], "")
	}
	if properties["key-without-value-3"] != "" {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", properties["key-without-value-3"], "")
	}
	if properties["key-with-value-1"] != "azerty" {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", properties["key-with-value-1"], "azerty")
	}
	if properties["key-with-value-2"] != "123 123 123 123" {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", properties["key-with-value-2"], "123 123 123 123")
	}
	if properties["key-with-value-3"] != "aze = aze = aze" {
		t.Errorf("ReadFromFileToMap() - error - got = '%v', want '%v'", properties["key-with-value-3"], "aze = aze = aze")
	}
}
