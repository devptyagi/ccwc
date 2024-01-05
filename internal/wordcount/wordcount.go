package wordcount

import (
	"bufio"
	"bytes"
)

// CountWords counts the number of words in the given input.
func CountWords(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanWords)
	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
