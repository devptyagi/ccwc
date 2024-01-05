package main

import (
	"bufio"
	"ccwc/internal/bytecount"
	"ccwc/internal/charactercount"
	"ccwc/internal/linecount"
	"ccwc/internal/wordcount"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func printUsage() {
	usage := `
Usage: ccwc [options] <file>

Options:
  -c    Count bytes in the file
  -l    Count lines in the file
  -m    Count characters in the file
  -w    Count words in the file
  --help    Show this help message

If no file is provided, ccwc reads from stdin.

If no option is specified, ccwc defaults to -l -w and -c flags all together.
`
	fmt.Println(usage)
}

func main() {
	// Define flags
	byteCountFlag := flag.Bool("c", false, "count bytes")
	lineCountFlag := flag.Bool("l", false, "count lines")
	charCountFlag := flag.Bool("m", false, "count characters")
	wordCountFlag := flag.Bool("w", false, "count words")

	// Custom usage function for --help
	flag.Usage = printUsage

	// Parse the flags
	flag.Parse()

	var content []byte
	var err error
	var fromFile bool

	// Check if a filename is provided
	if flag.NArg() == 1 {
		// Read the file
		filename := flag.Arg(0)
		fromFile = true
		content, err = os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error reading file: %s", err)
		}
	} else {
		// Read from stdin
		fromFile = false
		reader := bufio.NewReader(os.Stdin)
		content, err = io.ReadAll(reader)
		if err != nil {
			log.Fatalf("Error reading standard input: %s", err)
		}
	}

	printResult := func(output string) {
		if fromFile {
			fmt.Printf("%s %s\n", output, flag.Arg(0))
		} else {
			fmt.Printf("%s\n", output)
		}
	}

	// Perform counts based on flags
	anyFlagSet := *byteCountFlag || *lineCountFlag || *charCountFlag || *wordCountFlag

	if !anyFlagSet {
		byteCount := bytecount.CountBytes(content)
		lineCount := linecount.CountLines(content)
		wordCount := wordcount.CountWords(content)
		printResult(fmt.Sprintf("%d %d %d", lineCount, wordCount, byteCount))
	}

	if *byteCountFlag {
		count := bytecount.CountBytes(content)
		printResult(fmt.Sprintf("%d", count))
	}
	if *lineCountFlag {
		count := linecount.CountLines(content)
		printResult(fmt.Sprintf("%d", count))
	}
	if *charCountFlag {
		count := charactercount.CountCharacters(content)
		printResult(fmt.Sprintf("%d", count))
	}
	if *wordCountFlag {
		count := wordcount.CountWords(content)
		printResult(fmt.Sprintf("%d", count))
	}
}
