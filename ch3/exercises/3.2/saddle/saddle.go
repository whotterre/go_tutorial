package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 900, 700       // canvas size in pixels
	cells         = 100            // number of grid cells
	xyrange       = 30.0           // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.01   // pixels per z unit
	angle         = math.Pi / 6    // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	// Create the SVG output file
	file, err := os.Create("surface.svg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the SVG header
	_, err = fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	if err != nil {
		fmt.Println("Error writing SVG header:", err)
		return
	}

	// Generate the grid and write polygons to the file
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// Skip invalid polygons
			if isInvalid(ax, ay) || isInvalid(bx, by) || isInvalid(cx, cy) || isInvalid(dx, dy) {
				continue
			}

			_, err := fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			if err != nil {
				fmt.Println("Error writing polygon:", err)
				return
			}
		}
	}

	// Write the SVG footer
	_, err = fmt.Fprintln(file, "</svg>")
	if err != nil {
		fmt.Println("Error writing SVG footer:", err)
		return
	}

	fmt.Println("SVG file successfully generated: surface.svg")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	fmt.Println(y)
    return (math.Pow(x, 2) - math.Pow(y, 2))/ 15
}
// Check if a point is invalid
func isInvalid(x, y float64) bool {
	return math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0)
}
