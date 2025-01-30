package main 

// Define a struct for a concrete product (AK47) that implements the Gun interface

type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun : Gun{
			name: "AK47 Gun",
			power: 5,
		},
	}
}
