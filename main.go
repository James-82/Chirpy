package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	// Creating a new instance of http.ServeMux using the http.NewServeMux function
	mux := http.NewServeMux()

	// Registering a handler function with the ServeMux instance
	mux.HandleFunc("/", mainHandler)

	// Creating and configuring an HTTP server with the ServeMux instance as the handler
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// Starting the server
	fmt.Println("Starting server on :8080")
	server.ListenAndServe()
}
