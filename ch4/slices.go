package main

import "fmt"

func main() {
	// Slice literal
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[2:])
	// Prints out elements of s
	for _, item := range s {
		fmt.Printf("%d\n", item)
	}
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}