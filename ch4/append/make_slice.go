package main

import "fmt"

func main() {
	arr := [6]int{9,8,7,6,5,4}
	made := make([]int, 3, 5)

	for _, num := range arr {
		made = append(made, num)
	}
	// Prints elements of num
	for _, item := range made {
		fmt.Println(item)
	}
	fmt.Print("Cap ", cap(arr))
	fmt.Print("Len ", len(arr))

}