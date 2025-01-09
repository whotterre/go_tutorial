package main

import (
	"fmt"
)

func main() {
	fmt.Println("🚀 Initiating launch sequence...")
	defer missionControl()

	fmt.Println("🌌 Entering orbit calculations...")
	calculateOrbit()

	fmt.Println("This message won't appear due to a panic.")
}

func calculateOrbit() {
	// Simulate a critical error
	fmt.Println("💥 Critical error: Failed orbital parameters!")
	panic("Orbital calculations failed!")
}

func missionControl() {
	if r := recover(); r != nil {
		fmt.Println("🛑 Mission control to crew: We've encountered a critical error.")
		fmt.Printf("📡 Error details: %v\n", r)
		fmt.Println("📖 Taking corrective measures and reinitializing sequence...")
	}

	fmt.Println("✅ Mission control: All systems nominal. Proceed with caution.")
}
