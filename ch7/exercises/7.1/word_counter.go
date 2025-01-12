package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type Counter interface {
	Write(p []byte) (int, error)
}

// Counts number of words in the input
type WordCounter int

// Counts number of lines in the input
type LineCounter int

// Write method for WordCounter
func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count += 1
	}
	*wc += WordCounter(count)
	return len(p), nil
}

// Write method for LineCounter
func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count += 1
	}
	*lc += LineCounter(count)
	return len(p), nil
}

func main() {
	var wc WordCounter
	var lc LineCounter

	input := "Hello, world!\n"

	wc.Write([]byte(input))
	fmt.Printf("Word count: %d\n", wc)

	lc.Write([]byte(input))
	fmt.Printf("Line count: %d\n", lc)
}
