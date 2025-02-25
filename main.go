package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Extract the path after /welcome/
	name := strings.TrimPrefix(r.URL.Path, "/welcome/")

	// Log the request details
	log.Printf(
		"Request: %s %s | IP: %s | User-Agent: %s",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
		r.UserAgent(),
	)

	if name == "" {
		fmt.Fprint(w, "Anonymous")
	} else {
		fmt.Fprintf(w, "Welcome %s", name)
	}

	// Log processing time
	log.Printf("Processed request in %v", time.Since(start))
}

func main() {
	fmt.Println("Server listening on port 5000...")
	http.HandleFunc("/welcome/", welcomeHandler)
	http.ListenAndServe(":5000", nil)
}
