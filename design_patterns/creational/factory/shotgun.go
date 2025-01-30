package main

// Define a struct for a concrete product (Shotgun) that implements the Gun interface

type Shotgun struct {
	Gun
}

func newShotgun() IGun {
	return &Shotgun{
		Gun: Gun{
			name:  "Shotgun Gun",
			power: 15,
		},
	}
}
