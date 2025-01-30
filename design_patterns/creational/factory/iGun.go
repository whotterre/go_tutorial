package main

// Create a common interface, defining all methods all "guns" should have
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

