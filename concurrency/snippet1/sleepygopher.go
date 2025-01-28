package main 

import (
	"fmt"
	"time"
)

func main(){
	go sleepyGopher()
	time.Sleep(4 * time.Second)
}
// Waits 3 seconds then prints "snore message"
func sleepyGopher(){
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}