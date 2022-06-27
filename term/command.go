package term

import (
	"fmt"
	"log"
	"os"
)

var output = os.Stdout

// Clear screen
func Clear() {
	_, _ = fmt.Fprintf(output, "\033[2J")
}

// Move to line / column
func MoveTo(row int, column int) {
	_, _ = fmt.Fprintf(output, "\033[%d;%dH", row, column)
	log.Printf("") // FIXME
}
