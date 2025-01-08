package main

import (
	"ch5/exercises/5.2/helpers"
	"golang.org/x/net/html"
	"os"
	"log"
	"fmt"
)


func main()  {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	countMap := make(map[string]int)
	helpers.GetCountMap(doc, countMap)
	
	for tag, count := range countMap {
                fmt.Printf("%s: %d\n", tag, count)
        }

}
