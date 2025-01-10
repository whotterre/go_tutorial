package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func main() {
	p := Point{1, 2}
	q := Point{3, 4}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	

	// Method 1 of calling the ScaleBy fn
	r := &Point{1.0, 2.0}
	r.ScaleBy(3)
	fmt.Println(*r)
	// Method 2 of calling the ScaleBy fn
	j := Point{1, 2}
	pptr := &j
	pptr.ScaleBy(2)
	fmt.Println(j)
	// Method 3 of calling the ScaleBy fn
	t := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Print(t)

}

// Package level function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Method of p type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		fmt.Print(i)
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Methods with pointer reciever variable
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
