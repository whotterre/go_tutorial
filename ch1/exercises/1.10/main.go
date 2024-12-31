package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run fetchall.go <URLs>")
		return
	}

	start := time.Now()
	ch := make(chan string)
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		output := <-ch
		fmt.Println(output)
		_, err := outputFile.WriteString(output + "\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	// Add "http://" prefix if missing
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Write response to a file for inspection
	fileName := strings.Replace(url, "http://", "", 1)
	fileName = strings.Replace(fileName, "https://", "", 1)
	fileName = strings.Replace(fileName, "/", "_", -1) + ".txt"

	outputFile, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprintf("creating file for %s: %v", url, err)
		return
	}
	defer outputFile.Close()

	nbytes, err := io.Copy(outputFile, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
