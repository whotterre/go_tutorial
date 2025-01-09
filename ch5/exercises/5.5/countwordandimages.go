package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)
// CountWordsAndImages fetches a URL and returns the number of words and images in the document.
func CountWordsAndImages(url string) (words, images int, err error) {
	// Fetch the URL
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("fetching URL: %v", err)
	}
	defer resp.Body.Close()

	// Parse the HTML document
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing HTML: %v", err)
	}

	// Count words and images in the document
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages recursively traverses the HTML node tree and counts words and images.
func countWordsAndImages(n *html.Node) (words, images int) {
	// Check the node type
	if n.Type == html.TextNode {
		// Count words in text nodes
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		// Count <img> elements
		images++
	}

	// Recursively process child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	return
}

func main() {
	// Example URL
	url := "https://www.example.com"

	// Call CountWordsAndImages
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("Words: %d, Images: %d\n", words, images)
}
