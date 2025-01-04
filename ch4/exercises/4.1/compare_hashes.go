package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("8"))
	fmt.Println(compareHashes(a, b))
	
}
//Returns number of differing bits in two SHA256 hashes
func compareHashes(hashOne [32]byte, hashTwo[32]byte) int{
	count := 0
	for i := 0; i < 32; i++ {
		diff := hashOne[i] ^ hashTwo[i]
		count += popCount(diff)
		}
	return count
}
//Returns number of set bits in a byte
func popCount(b byte) int {
	count := 0
	for b > 0 {
		count += int(b & 1)
		b >>= 1
	}
	return count
}