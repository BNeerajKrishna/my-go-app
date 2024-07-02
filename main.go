package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// API routes
	hostname, err := os.Hostname()

	// Serve files from static folder
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Serve api /hi
	http.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, hostname is %v and is error is %v\n", hostname, err)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port )

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}