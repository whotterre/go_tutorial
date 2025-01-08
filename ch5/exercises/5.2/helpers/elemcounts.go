package helpers

import (
	"golang.org/x/net/html"
)


// Gets count of elements in a html document
func GetCountMap(n *html.Node, countMap map[string]int){
	if n == nil {
		return 
	}
	// Check if the element's name does not exist in the count map
	if n.Type == html.ElementNode {
		countMap[n.Data] += 1
	}

	GetCountMap(n.FirstChild, countMap)
	GetCountMap(n.NextSibling, countMap)
	
}
