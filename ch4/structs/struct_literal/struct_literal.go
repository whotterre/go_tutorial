package main

import "fmt"

func main() {
	// Empty struct
	type Empty struct{}

	// Form 1 of a struct literal
	type Point struct{ X, Y int }
	p := Point{1, 2}

	// Form 2 of a struct literal
	type CoOrd struct{ X, Y int }
	r := CoOrd{X: 3, Y: 9}

	// Structs as function arguments
	Scale := func(p Point, factor int) Point {
		return Point{p.X * factor, p.Y * factor}
	}

	// Example usage of Scale
	scaledPoint := Scale(p, 2)
	fmt.Println("Original Point:", p)
	fmt.Println("Scaled Point:", scaledPoint)
	fmt.Println("Another Struct Literal:", r)
}
