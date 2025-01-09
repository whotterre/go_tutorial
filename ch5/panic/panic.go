package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Draw a card and check its suit
	s := drawCard()
	switch s {
	case "Spades":
		fmt.Println("Shades")
	case "Hearts":
		fmt.Println("Love")
	case "Diamonds":
		fmt.Println("Wealth")
	case "Clubs":
		fmt.Println("Growth")
	default:
		panic(fmt.Sprintf("Invalid suit %q", s))
	}

}

func drawCard() string {
	// Define the suits
	cards := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	// Pick a random card
	return cards[rand.Intn(len(cards))]
}

