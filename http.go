package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	hostname, _ := os.Hostname()
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// Empty handler for /favicon.ico by browser
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().String()
		fmt.Fprintf(os.Stdout, "I'm %s time now: %s\n", hostname, currentTime)
		fmt.Fprintf(w, "I'm %s time now: %s\n", hostname, currentTime)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
