package main
// Solution to exercise 5.1
import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/net/html"
)

func main() {
	// Parse the html to derive a DOM tree 
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	findLinks(doc, 0)

	


}

// Finds links in <a> tags in a html document
func findLinks(n *html.Node, depth int) {
	if n == nil {
		return 
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("%s%s\n", indent(depth), a.Val)
			}
		}
	}
	findLinks(n.FirstChild, depth + 1)
	findLinks(n.NextSibling, depth)
}

func indent(depth int) string{
	return strings.Repeat(" ", depth)
}
