package main

import (
	"fmt"
)

func main() {
	hashmap := make(map[string]int)

	hashmap["Java"] = 1
	hashmap["TypeScript"] = 5
	hashmap["Go"] = 4
	hashmap["Go"] = 2
	for key, value := range hashmap {
		fmt.Println(key, value)
	}
	// Enumerate for key and presence boolean
	item, ok := hashmap["Java"]
	if ok {
		fmt.Print("Java in hashmap with val ", item)
	} else {
		fmt.Print("Java not found")
	}
	
}
