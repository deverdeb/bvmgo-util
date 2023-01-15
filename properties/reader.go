package properties

import (
	"bytes"
	"fmt"
	"github.com/deverdeb/bvmgo-util/errors"
	"github.com/deverdeb/bvmgo-util/iostring"
	"io"
	"os"
	"strings"
)

func ReadFromFileToMap(filename string) (map[string]string, error) {
	propFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.NewWithCause(err, "failed to open '%s' properties file", filename)
	}
	defer closeFile(propFile)
	properties, err := ReadToMap(propFile)
	return properties, err
}

func ReadFromBytesToMap(content []byte) (map[string]string, error) {
	reader := bytes.NewReader(content)
	return ReadToMap(reader)
}

func ReadToMap(reader io.Reader) (map[string]string, error) {
	result := make(map[string]string)
	lineReader := iostring.NewLineReaderFromReader(reader)
	lineNumber := 0
	for lineReader.HasNext() {
		lineNumber++
		key, value, isValue, err := extractDataFromLine(lineReader)
		if err != nil {
			return result, err
		}
		if isValue {
			result[key] = value
		}
	}
	return result, nil
}

func extractDataFromLine(lineReader *iostring.LineReader) (key string, value string, isValue bool, err error) {
	lineContent, lineNumber, isOk := lineReader.Read()
	if !isOk {
		return "", "", false, fmt.Errorf("failed to read data at line %d", lineNumber)
	}
	content := strings.TrimSpace(lineContent)
	if content == "" || content[0:0] == "#" {
		// Empty or comment line - ignore
		return "", content, false, nil
	} else {
		equalIndex := strings.IndexRune(lineContent, '=')
		if equalIndex < 0 {
			// Key without value
			key = lineContent
		} else {
			key = strings.TrimSpace(lineContent[0:equalIndex])
			value = strings.TrimSpace(lineContent[equalIndex+1:])
		}
		return key, value, true, nil
	}
}

func closeFile(file *os.File) {
	if file != nil {
		_ = file.Close()
	}
}
