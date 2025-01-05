package main

import (
	//"fmt"
	"fmt"
	"unicode"
)

func unicodeToAsciiSpace(slice []byte){
	// Squashes each run of adjacent Unicode spaces(see unicode.IsSpace) in
	// a UTF-8-encoded []byte slice into a single ASCII space

	writeIdx := 0

	prevWasSpace := false

	for _, b := range slice {
		if unicode.IsSpace(rune(b)){
			//If prev character was a space
			if prevWasSpace {
				continue
			}
			// Append an ASCII space string and set flag to true
			slice[writeIdx] = ' '
			writeIdx++
			prevWasSpace = true
		} else {

			slice[writeIdx] = b
			writeIdx++
			prevWasSpace = false
		}
	}
}
func main(){
	str := "This is a byte"
	unicodeToAsciiSpace([]byte(str))
	fmt.Print(str)
}