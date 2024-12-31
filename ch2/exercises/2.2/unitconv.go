package unitconv

import (
	"bufio"
	"fmt"
	"os"
)

// Conversion functions for different units

// Celsius to Fahrenheit
func CToF(celsius float64) float64 {
	return celsius*9/5 + 32
}

// Fahrenheit to Celsius
func FToC(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Meters to Feet
func MToF(meters float64) float64 {
	return meters * 3.28084
}

// Feet to Meters
func FToM(feet float64) float64 {
	return feet * 0.3048
}

// Kilograms to Pounds
func KToP(kilograms float64) float64 {
	return kilograms * 2.20462
}

// Pounds to Kilograms
func PToK(pounds float64) float64 {
	return pounds * 0.453592
}


func main() {
	
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			handleConversion(arg)
		}
	} else {
		// If no arguments, read from standard input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter conversions in the format: <value> <unit>")
		for scanner.Scan() {
			handleConversion(scanner.Text())
		}
	}
}
