package main

import (
	"fmt"
	"time"
)

func main() {
	// Creates a channel 
	c := make(chan int)
	// The time.After returns a timeout if the time arg is exceeded without recieving a value
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		// If a goroutine sends a value to the channel after 2s, print that
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, "has finished sleeping")
		case <-timeout:
			// Otherwise, do this...
			fmt.Println("couldn't wait anymore")
			return 
		}
	}

}
