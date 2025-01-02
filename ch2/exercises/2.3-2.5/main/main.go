package main

import (
	"fmt"
	"time"

	"2.3-2.5/popcount"
)

func main() {
	// Times function that finds number of set bits with bitwise operations
	
	start := time.Now()
	popcount.PopCount(43)
	end := time.Since(start)

	begin := time.Now()
	popcount.LoopCount(43)
	term := time.Since(begin)

	fmt.Printf("Bitwise popcount finished in %v\n", end)
	fmt.Printf("Loopwise popcount finished in %v\n", term)

}