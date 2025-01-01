package main

import (
	"fmt"
	"time"

	"2.3/popcount"
)

func main() {
	// Times function that finds number of set bits with bitwise operations
	
	start := time.Now()
	popcount.PopCount(43)
	end := time.Since(start)

	fmt.Printf("Bitwise popcount finished in %v", end)
	

}