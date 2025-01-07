package main

import (
	"ch4/exercises/utils"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	res, err := utils.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var (
		lessThanMonth []utils.Issue
		lessThanYear  []utils.Issue
		moreThanYear  []utils.Issue
	)

	now := time.Now()
	for _, item := range res.Items {
		age := now.Sub(item.CreatedAt)
		switch {
		case age < (30 * 24 * time.Hour):
			lessThanMonth = append(lessThanMonth, *item)
		case age < (365 * 24 * time.Hour):
			lessThanYear = append(lessThanYear, *item)
		default:
			moreThanYear = append(moreThanYear, *item)
	}

	// Print results
	fmt.Println("Issues less than a month old:")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nIssues less than a year old:")
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nIssues more than a year old:")
	for _, item := range moreThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

}