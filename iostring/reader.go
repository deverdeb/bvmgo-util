package iostring

import (
	"bufio"
	"bytes"
	"io"
)

// LineReader reads text line by line.
type LineReader struct {
	// scanner is used to read lines
	scanner *bufio.Scanner
	// lineNumber is the current line number
	lineNumber int
	// hasNext line ?
	hasNext bool
}

// NewLineReaderFromBytes creates a line reader from a byte array.
func NewLineReaderFromBytes(content []byte) *LineReader {
	byteReader := bytes.NewReader(content)
	return NewLineReaderFromReader(byteReader)
}

// NewLineReaderFromReader creates a line reader from a byte reader.
func NewLineReaderFromReader(reader io.Reader) *LineReader {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	hasNext := scanner.Scan()
	/*if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}*/
	lineReader := &LineReader{
		scanner:    scanner,
		lineNumber: 0,
		hasNext:    hasNext,
	}
	return lineReader
}

// Read returns next text line and line number.
// Return isOk=false and an empty line at end of data stream (if reader can not read line).
func (reader *LineReader) Read() (lineContent string, lineNumber int, isOk bool) {
	if reader.hasNext {
		reader.lineNumber++
		content := reader.scanner.Text()
		reader.hasNext = reader.scanner.Scan()
		return content, reader.lineNumber, true
	}
	return "", reader.lineNumber, false
}

// HasNext returns if reader can read other text lines.
func (reader *LineReader) HasNext() bool {
	return reader.hasNext
}

// LineNumber returns current line position (0 = not start).
func (reader *LineReader) LineNumber() int {
	return reader.lineNumber
}
