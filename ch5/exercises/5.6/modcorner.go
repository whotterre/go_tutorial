package main

import "math"

const (
	xyrange = 30
	xyscale = 40
	cells   = 50
	width   = 300
	height  = 500
	angle   = 30
	zscale  = 2
)

func corner(i, j int) (sx, sy float64) {
	sin30, cos30 := math.Sin(angle), math.Cos(angle)
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale

	// Use a bare return statement.
	return
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
