package iostring

import (
	"os"
	"testing"
)

var textTest = `line1
line 2
  line 3  
`

func TestLineReader_NewLineReaderFromBytes(t *testing.T) {
	bytes := []byte(textTest)
	reader := NewLineReaderFromBytes(bytes)
	if !reader.HasNext() {
		t.Errorf("LineReader.HasNext() - 1 = %v, want %v", reader.HasNext(), true)
		return
	}
	line1, num, ok := reader.Read()
	if !ok {
		t.Errorf("LineReader.Read() - ok = '%v', want '%v'", ok, true)
	}
	if num != 1 {
		t.Errorf("LineReader.Read() - line number = '%v', want '%v'", num, 1)
	}
	if line1 != "line1" {
		t.Errorf("LineReader.HasNext() - value = '%v', want '%v'", line1, "line1")
	}

	if !reader.HasNext() {
		t.Errorf("LineReader.HasNext() - 2 = %v, want %v", reader.HasNext(), true)
		return
	}
	line2, num, ok := reader.Read()
	if !ok {
		t.Errorf("LineReader.Read() - ok = '%v', want '%v'", ok, true)
	}
	if num != 2 {
		t.Errorf("LineReader.Read() - line number = '%v', want '%v'", num, 2)
	}
	if line2 != "line 2" {
		t.Errorf("LineReader.HasNext() - value = '%v', want '%v'", line2, "line 2")
	}

	if !reader.HasNext() {
		t.Errorf("LineReader.HasNext() - 3 = %v, want %v", reader.HasNext(), true)
		return
	}
	line3, num, ok := reader.Read()
	if !ok {
		t.Errorf("LineReader.Read() - ok = '%v', want '%v'", ok, true)
	}
	if num != 3 {
		t.Errorf("LineReader.Read() - line number = '%v', want '%v'", num, 3)
	}
	if line3 != "  line 3  " {
		t.Errorf("LineReader.HasNext() - value = '%v', want '%v'", line3, "  line 3  ")
	}

	if reader.HasNext() {
		t.Errorf("LineReader.HasNext() - end = %v, want %v", reader.HasNext(), false)
	}
}

func TestLineReader_NewLineReaderFromReader(t *testing.T) {
	propFile, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("TestLineReader_NewLineReaderFromBytes - failed to read test file 'test.txt'")
		return
	}
	defer propFile.Close()
	reader := NewLineReaderFromReader(propFile)
	if !reader.HasNext() {
		t.Errorf("LineReader.HasNext() - 1 = %v, want %v", reader.HasNext(), true)
		return
	}
	line1, num, ok := reader.Read()
	if !ok {
		t.Errorf("LineReader.Read() - ok = '%v', want '%v'", ok, true)
	}
	if num != 1 {
		t.Errorf("LineReader.Read() - line number = '%v', want '%v'", num, 1)
	}
	if line1 != "line1" {
		t.Errorf("LineReader.HasNext() - value = '%v', want '%v'", line1, "line1")
	}

	if !reader.HasNext() {
		t.Errorf("LineReader.HasNext() - 2 = %v, want %v", reader.HasNext(), true)
		return
	}
	line2, num, ok := reader.Read()
	if !ok {
		t.Errorf("LineReader.Read() - ok = '%v', want '%v'", ok, true)
	}
	if num != 2 {
		t.Errorf("LineReader.Read() - line number = '%v', want '%v'", num, 2)
	}
	if line2 != "line 2" {
		t.Errorf("LineReader.HasNext() - value = '%v', want '%v'", line2, "line 2")
	}

	if !reader.HasNext() {
		t.Errorf("LineReader.HasNext() - 3 = %v, want %v", reader.HasNext(), true)
		return
	}
	line3, num, ok := reader.Read()
	if !ok {
		t.Errorf("LineReader.Read() - ok = '%v', want '%v'", ok, true)
	}
	if num != 3 {
		t.Errorf("LineReader.Read() - line number = '%v', want '%v'", num, 3)
	}
	if line3 != "  line 3" {
		t.Errorf("LineReader.HasNext() - value = '%v', want '%v'", line3, "  line 3")
	}

	if reader.HasNext() {
		t.Errorf("LineReader.HasNext() - end = %v, want %v", reader.HasNext(), false)
	}
}
