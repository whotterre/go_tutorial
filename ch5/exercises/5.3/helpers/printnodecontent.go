package helpers

import (
	"fmt"
	"golang.org/x/net/html"
)

// PrintNodeContent prints the text content of all text nodes in an HTML document.
// It skips <script> and <style> elements.
func PrintNodeContent(n *html.Node) {
	if n == nil {
		return
	}

	// If it's a <script> or <style> element, skip its children
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	// If it's a text node, print its content
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	// Recursively process child and sibling nodes
	PrintNodeContent(n.FirstChild)
	PrintNodeContent(n.NextSibling)
}
