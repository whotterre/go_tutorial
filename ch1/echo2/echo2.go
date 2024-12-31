package main

import (
	"fmt"
	"os"
)
//Input via command line args via for loop variation 2
func main(){
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}