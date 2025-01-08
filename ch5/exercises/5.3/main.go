package main

import (
	"ch5/exercises/5.3/helpers"
	"golang.org/x/net/html"
	"os"
	"log"
)


func main()  {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	helpers.PrintNodeContent(doc)
	

}
