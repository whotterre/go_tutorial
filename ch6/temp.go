package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
}

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, log float64
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("%v", t.average())
	fmt.Printf("%v", report.temperature.average())
}
