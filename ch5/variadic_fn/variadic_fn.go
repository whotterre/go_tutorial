package main

import "fmt"

func main(){
	addAll(2,28,3,3,3,43)
}
// Variadic function that takes variable args
func addAll(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	fmt.Println(nums)
	return total
}
