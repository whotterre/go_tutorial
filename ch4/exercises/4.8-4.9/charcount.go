package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Counts letters, digits and so on in their Unicode categories
func CharCount() {	
	counts := make(map[rune]int)
	letters := 0
	digits := 0
	whitespace := 0
	punctuation := 0
	invalid := 0
	var utflen [utf8.UTFMax + 1]int

	in := bufio.NewReader(os.Stdin)
	// Infinite loop that awaits input
	for {
		// Rune, number of bytes read, errors (if any)
		r, n, err := in.ReadRune()
		// Check if we're at the end of the file
		if err == io.EOF {
			break
		}
		/* If an error occurs, exit prematurely
		and print character count
		*/
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		} 
		if unicode.IsLetter(r){
			letters++
		}
		if unicode.IsDigit(r){
			digits++
		} else if unicode.IsPunct(r){
			punctuation++
		} else if unicode.IsSpace(r) {
			whitespace++
		}
		//Update for all runes counted 
		counts[r]++
		//Update UTF-8 byte length counts
		utflen[n]++
	}
	fmt.Printf("Letters: %d\n", letters)
	fmt.Printf("Digits: %d\n", digits)
	fmt.Printf("Punctuation: %d\n", punctuation)
	fmt.Printf("Whitespace: %d\n", whitespace)
	fmt.Printf("Invalid characters: %d\n", invalid)

	fmt.Println("Unicode Category Counts:")
	for c, count := range counts {
		fmt.Printf("%q: %d\n", c, count)
	}

	// Print byte length frequencies
	fmt.Println("UTF-8 Byte Lengths:")
	for i, count := range utflen {
		if count > 0 {
			fmt.Printf("%d-byte characters: %d\n", i, count)
		}
	}
}
//Writes the output of char count to a file
func WordFreq(){
	file, err := os.Create("utfstats.txt")

	if err != nil {
		fmt.Print("Error writing to file")
		os.Exit(1)
	}
	defer file.Close()
	
	counts := make(map[rune]int)
	letters := 0
	digits := 0
	whitespace := 0
	punctuation := 0
	invalid := 0
	var utflen [utf8.UTFMax + 1]int

	in := bufio.NewReader(os.Stdin)
	// Infinite loop that awaits input
	for {
		// Rune, number of bytes read, errors (if any)
		r, n, err := in.ReadRune()
		// Check if we're at the end of the file
		if err == io.EOF {
			break
		}
		/* If an error occurs, exit prematurely
		and print character count
		*/
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		} 
		if unicode.IsLetter(r){
			letters++
		}
		if unicode.IsDigit(r){
			digits++
		} else if unicode.IsPunct(r){
			punctuation++
		} else if unicode.IsSpace(r) {
			whitespace++
		}
		//Update for all runes counted 
		counts[r]++
		//Update UTF-8 byte length counts
		utflen[n]++
	}
	fmt.Fprintf(file, "Letters: %d\n", letters)
	fmt.Fprintf(file, "Digits: %d\n", digits)
	fmt.Fprintf(file, "Punctuation: %d\n", punctuation)
	fmt.Fprintf(file, "Whitespace: %d\n", whitespace)
	fmt.Fprintf(file, "Invalid characters: %d\n", invalid)

	fmt.Fprintf(file, "Unicode Category Counts:")
	for c, count := range counts {
		fmt.Fprintf(file, "%q: %d\n", c, count)
	}

	// Print byte length frequencies
	fmt.Fprintf(file, "UTF-8 Byte Lengths:")
	for i, count := range utflen {
		if count > 0 {
			fmt.Fprintf(file, "%d-byte characters: %d\n", i, count)
		}
	}
}