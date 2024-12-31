package main

import (
	"fmt"
	"os"
)
//Input via command line args via for loop variation 1
func main(){
	var s, sep string 
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}