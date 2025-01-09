package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "golang.org/x/net/html"
)

func main() {
    const url = "http://www.bing.com"

    // Method 1: Error propagation
    res, err := http.Get(url)
    if err != nil {
        log.Fatalf("Failed to fetch URL %s: %v", url, err)
    }
    defer res.Body.Close()

    doc, err := html.Parse(res.Body)
    if err != nil {
        log.Fatalf("Failed to parse page at %s as HTML: %v", url, err)
        fmt.Print(doc)
    }

    // Method 2: Delayed retries
    if err := waitForServer(url); err != nil {
        log.Fatalf("Site is down: %v", err)
    }
}

// waitForServer attempts to connect to the server with a delay between retries.
func waitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil
        }
        log.Printf("Server not responding (%s), retrying...", url)
        time.Sleep(time.Duration(1<<uint(tries)) * time.Second)
    }
    return fmt.Errorf("server %s failed to respond after %v", url, timeout)
}
