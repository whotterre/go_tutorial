package main

import (
	"ch4/exercises/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <API_KEY> <MOVIE_NAME>")
		os.Exit(1)
	}

	apiKey := os.Args[1]
	movieName := strings.Join(os.Args[2:], " ")

	// Call the function to download the movie poster
	posterFileName, err := utils.GetPosterImage(apiKey, movieName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Poster downloaded successfully: %s\n", posterFileName)
}
