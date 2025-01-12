package main

import (
	"fmt"

	"7.1/counters"
)



func main() {
	var wc counter.WordCounter
	var lc counter.LineCounter

	input := "Hello, world!\n"

	

	wc.Write([]byte(input))
	fmt.Printf("Word count: %d\n", wc)

	lc.Write([]byte(input))
	fmt.Printf("Line count: %d\n", lc)
}
