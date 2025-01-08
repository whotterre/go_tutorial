package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// 4.10
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//Ex 4.13

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
	Error  string `json:"Error"`
}

func GetPosterImage(apiKey, movieName string) (string, error) {
	//Format the query string
	queryURL := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", movieName, apiKey)
	res, err := http.Get(queryURL)
	//Handle error in GET Request
	if err != nil {
		return "", fmt.Errorf("error fetching movie details: %v", err)
	}
	// Dealloc mem for res.Body after use
	defer res.Body.Close()
	//Handle if status code isn't 200
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: HTTP %d", res.StatusCode)
	}
	var movie Movie
	if err := json.NewDecoder(res.Body).Decode(&movie); err != nil {
		return "", fmt.Errorf("Error decoding JSON response: %v", err)
	}
	// Handle cases where there is an error and there is no poster image
	if movie.Error != "" {
		return "", errors.New(movie.Error)
	}
	if movie.Poster == "N/A" {
		return "", errors.New("No poster available for this movie")
	}
	// Download poster
	posterResp, err := http.Get(movie.Poster)
	if err != nil {
		return "", fmt.Errorf("Error downloading poster: %v", err)
	}
	defer posterResp.Body.Close()
	if posterResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: HTTP %d when downloading poster", posterResp.StatusCode)
	}

	posterFileName := strings.ReplaceAll(movie.Title, " ", "_") + ".jpg"
	file, err := os.Create(posterFileName)

	if err != nil {
		return "", fmt.Errorf("Error creating file")
	}
	defer file.Close()

	// Write the poster data to the file
	_, err = io.Copy(file, posterResp.Body)
	if err != nil {
		return "", fmt.Errorf("error saving poster: %v", err)
	}

	return posterFileName, nil

}
