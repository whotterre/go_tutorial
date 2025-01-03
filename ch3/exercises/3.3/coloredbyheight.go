package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.1
	angle         = math.Pi / 8
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	file, err := os.Create("surface_colored.svg")
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

	// Find min and max z-values
	zmin, zmax := findHeightRange()

	// Generate the grid and write polygons to the file
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if isInvalid(ax, ay) || isInvalid(bx, by) || isInvalid(cx, cy) || isInvalid(dx, dy) {
				continue
			}

			// Get average height for the polygon
			avgHeight := averageHeight(i, j)

			// Map the height to a color
			color := heightToColor(avgHeight, zmin, zmax)

			// Write the polygon with the fill color
			_, err := fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='fill: %s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
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

	fmt.Println("Colored SVG file successfully generated: surface_colored.svg")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	return math.Sin(x) * math.Cos(y)
}

// Find min and max height (z-values)
func findHeightRange() (float64, float64) {
	zmin, zmax := math.Inf(1), math.Inf(-1)
	for i := 0; i <= cells; i++ {
		for j := 0; j <= cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if z < zmin {
				zmin = z
			}
			if z > zmax {
				zmax = z
			}
		}
	}
	return zmin, zmax
}

// Get the average height for a grid cell
func averageHeight(i, j int) float64 {
	x1 := xyrange * (float64(i)/cells - 0.5)
	y1 := xyrange * (float64(j)/cells - 0.5)
	x2 := xyrange * (float64(i+1)/cells - 0.5)
	y2 := xyrange * (float64(j+1)/cells - 0.5)
	z1 := f(x1, y1)
	z2 := f(x2, y1)
	z3 := f(x1, y2)
	z4 := f(x2, y2)
	return (z1 + z2 + z3 + z4) / 4
}

// Map height to a color
func heightToColor(z, zmin, zmax float64) string {
	normalized := (z - zmin) / (zmax - zmin) // Normalize to [0, 1]
	red := int(255 * normalized)            // Higher z -> more red
	blue := int(255 * (1 - normalized))     // Lower z -> more blue
	return fmt.Sprintf("#%02x00%02x", red, blue)
}

// Check if a point is invalid
func isInvalid(x, y float64) bool {
	return math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0)
}
