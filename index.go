package main

import (
	"fmt"
	"net/http"
)

// handler handles HTTP requests and sends a response.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Vercel! Your request path is: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	// Start the server on port 8080 (required by Vercel's serverless runtime).
	http.ListenAndServe(":8080", nil)
}
