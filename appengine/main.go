package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	serverHost = "libproxy-db.org"
	proxyURL   = "https://libproxy-db.org/proxies.json"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Set up a specific handler for GET requests for proxies.json.  The browser
	// extensions won't allow a redirect to a host that isn't in their security
	// policy, so until they are updated, we need to proxy it here.
	http.HandleFunc("/proxies.json", proxyHandler)

	http.HandleFunc("/", rootHandler)

	log.Println("listening on port", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rurl := *r.URL
	rurl.Scheme = "https"
	rurl.Host = serverHost

	http.Redirect(w, r, rurl.String(), http.StatusMovedPermanently)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	pr, err := http.NewRequestWithContext(r.Context(), http.MethodGet, proxyURL, nil)
	if err != nil {
		log.Printf("failed to create request: %v", err)
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(pr)
	if err != nil {
		log.Printf("failed to make proxy request: %v", err)
		http.Error(w, "failed to make request", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}
