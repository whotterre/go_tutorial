package main

import (
	"fmt"
	"math"
	"net/http"
)

var xyrange = 30.0 

func GenShape(w http.ResponseWriter, width, height int, color string) {
	var (
		cells    = 100 // Number of grid cells
		xyscale  = float64(width) / 2 / xyrange
		angle    = math.Pi / 8 
		sinTheta = math.Sin(angle)
		cosTheta = math.Cos(angle)
		zscale   = float64(height) * 0.1
	)

	// Write the SVG header
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", color, width, height)

	// Generate grid and write polygons
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, xyrange, xyscale, zscale, sinTheta, cosTheta, width, height)
			bx, by := corner(i, j, xyrange, xyscale, zscale, sinTheta, cosTheta, width, height)
			cx, cy := corner(i, j+1, xyrange, xyscale, zscale, sinTheta, cosTheta, width, height)
			dx, dy := corner(i+1, j+1, xyrange, xyscale, zscale, sinTheta, cosTheta, width, height)

			if isInvalid(ax, ay) || isInvalid(bx, by) || isInvalid(cx, cy) || isInvalid(dx, dy) {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	// Write the SVG footer
	fmt.Fprintln(w, "</svg>")
}

// Calculates the coordinates of a corner of the grid cell.
func corner(i, j int, xyrange, xyscale, zscale, sinTheta, cosTheta float64, width, height int) (float64, float64) {
	x := xyrange * (float64(i)/100.0 - 0.5)
	y := xyrange * (float64(j)/100.0 - 0.5)
	z := f(x, y)
	sx := float64(width)/2 + (x-y)*cosTheta*xyscale
	sy := float64(height)/2 + (x+y)*sinTheta*xyscale - z*zscale
	return sx, sy
}

// Validates if a coordinate is invalid.
func isInvalid(x, y float64) bool {
	return math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0)
}

// Calculates the height of the surface.
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return math.NaN()
	}
	return math.Sin(r) / r
}
