package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/img", handler)
	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

// HTTP handler to process query parameters and generate SVG.
func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Parse query parameters with defaults
	width := parseQueryParam(query.Get("width"), 300)
	height := parseQueryParam(query.Get("height"), 200)
	color := parseColor(query.Get("color"), "red")

	// Set response content type
	w.Header().Set("Content-Type", "image/svg+xml")

	// Generate SVG and write to response
	GenShape(w, width, height, color)
}

// Helper to parse integer query parameters with a default value
func parseQueryParam(param string, defaultValue int) int {
	if param == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(param)
	if err != nil {
		return defaultValue
	}
	return value
}

// Helper to normalize color parameter
func parseColor(color string, defaultColor string) string {
	if color == "" {
		return defaultColor
	}
	return strings.ToLower(color)
}
