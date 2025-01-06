package main

import "fmt"

func main() {
	// Define the composite structs
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Center Point
		Radius int
	}
	type Wheel struct {
		Circle Circle
		Spokes int
	}

	// Initialize a Wheel instance
	var w Wheel
	w.Circle.Center = Point{X: 8, Y: 8}
	w.Circle.Radius = 5
	w.Spokes = 20

	// Print the wheel structure
	fmt.Printf("Wheel: %+v\n", w)
}
