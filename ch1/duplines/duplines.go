package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Define map
	counts := make(map[string]int)
	//Get input (like Java Scanner)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
