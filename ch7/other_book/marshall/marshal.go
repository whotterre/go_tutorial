package main

import (
	"encoding/json"
	"fmt"
)

// Define the coords struct
type coords struct {
	Decimal     float64 `json:"decimal"`
	DMS         float64 `json:"dms"`
	Degrees     float64 `json:"degrees"`
	Minutes     int     `json:"minutes"`
	Seconds     int     `json:"seconds"`
	Hemisphere  string  `json:"hemisphere"`
}

// Implement the MarshalJSON method
func (c coords) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"decimal":    c.Decimal,
		"dms":        c.DMS,
		"degrees":    c.Degrees,
		"minutes":    c.Minutes,
		"seconds":    c.Seconds,
		"hemisphere": c.Hemisphere,
	}

	
	return json.Marshal(data)
}

func main() {
	// Create a coords instance
	coordinate := coords{
		Decimal:    37.7749,
		DMS:        37.7749,
		Degrees:    37,
		Minutes:    46,
		Seconds:    29,
		Hemisphere: "N",
	}

	// Marshal the coords struct into JSON
	jsonData, err := json.Marshal(coordinate)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON output
	fmt.Println(string(jsonData))
}
