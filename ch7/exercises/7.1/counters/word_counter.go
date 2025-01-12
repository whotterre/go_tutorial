package counter

import (
	"bufio"
	"bytes"
)

type Counter interface {
	Write(p []byte) (int, error)
}

// Counts number of words in the input
type WordCounter int


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

