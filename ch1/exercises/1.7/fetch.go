package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	/*
		The function call io.Copy(dst, src) reads from src and writes to dst. Use it
		instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
		buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
	*/
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Replace ioutil.ReadAll with io.Copy
		b, err := io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
}
