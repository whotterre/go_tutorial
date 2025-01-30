package main

/* A struct representing a product of the factory and methods to implement the factory
*/
type Gun struct {
	name string
	power int
}

func (g *Gun) setName(name string){
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int ){
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}



