package main

import (
	"fmt"
	"strings"
)

// Source of input into the pipeline
func sourceGopher(downstream chan string) {
	for _, v := range []string{"Hello Jenna!", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

// Filters "bad" gophers
func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			close(downstream)
			return
		}

		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

// Prints gophers from the upstream channel
func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	cOne := make(chan string)
	cTwo := make(chan string)

	go sourceGopher(cOne)
	go filterGopher(cOne, cTwo)
	printGopher(cTwo)
}
