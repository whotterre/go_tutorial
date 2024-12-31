package main

import (
	"fmt"
	"os"
)

func main(){
	for idx, val := range os.Args {
		fmt.Printf("%d, %s\n", idx, val)
	}
}