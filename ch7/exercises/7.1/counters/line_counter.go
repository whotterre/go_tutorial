package counter

import (
	"bufio"
	"bytes"
)

// Counts number of lines in the input
type LineCounter int

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