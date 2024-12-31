package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 255},     // Red
	color.RGBA{0, 255, 0, 255},     // Green
	color.RGBA{0, 0, 255, 255},     // Blue
	color.RGBA{255, 255, 0, 255},   // Yellow
	color.RGBA{0, 255, 255, 255},   // Cyan
	color.RGBA{255, 0, 255, 255},   // Magenta
	color.RGBA{128, 0, 128, 255},   // Purple
	color.RGBA{255, 165, 0, 255},   // Orange
	color.RGBA{255, 192, 203, 255}, // Pink
}

const (
	whiteIndex   = 0
	blackIndex   = 1
	redIndex     = 2
	greenIndex   = 3
	blueIndex    = 4
	yellowIndex  = 5
	cyanIndex    = 6
	magentaIndex = 7
	purpleIndex  = 8
	orangeIndex  = 9
	pinkIndex    = 10
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10    // Number of complete x oscillator cycles
		res     = 0.001 // Angular resolution
		size    = 200   // Image canvas size [-size..+size]
		nframes = 64    // Number of animation frames
		delay   = 8     // Delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // Relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase difference

	// Use colors from the palette for each point
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Choose color based on position (you can adjust the logic)
			colorIndex := (int(x*size+size) + int(y*size+size)) % len(palette)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Encode the GIF
}
