package linecount

import (
	"bufio"
	"bytes"
)

// CountLines counts the number of lines in the given input.
func CountLines(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
