package unitconv

import (
	"fmt"
	"strconv"
	"strings"
)

// Handle user input for conversion
func handleConversion(input string) {
	// Split input into value and unit
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Invalid input format")
		return
	}

	// Parse the value to float
	value, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return
	}

	unit := strings.ToLower(parts[1])

	// Perform conversion based on the unit
	switch unit {
	case "celsius":
		// Convert Celsius to Fahrenheit
		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", value, CToF(value))
	case "fahrenheit":
		// Convert Fahrenheit to Celsius
		fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", value, FToC(value))
	case "meters":
		// Convert Meters to Feet
		fmt.Printf("%.2f Meters = %.2f Feet\n", value, MToF(value))
	case "feet":
		// Convert Feet to Meters
		fmt.Printf("%.2f Feet = %.2f Meters\n", value, FToM(value))
	case "kilograms":
		// Convert Kilograms to Pounds
		fmt.Printf("%.2f Kilograms = %.2f Pounds\n", value, KToP(value))
	case "pounds":
		// Convert Pounds to Kilograms
		fmt.Printf("%.2f Pounds = %.2f Kilograms\n", value, PToK(value))
	default:
		fmt.Println("Unsupported unit:", unit)
	}
}
