package main

import (
	"fmt"
	"time"
)
// Goroutines are executed in random order 
func main(){
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher(id int){
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}