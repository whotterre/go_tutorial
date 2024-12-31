package main

import (
	//"container/heap"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	x := 1
	p := &x

	fmt.Println(*p)

	*p = 98
	*p++
	fmt.Println(x)
}