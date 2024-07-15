package main

import (
    "net/http"
)

// Custom handler that always returns a 404 status
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}

func main() {
    // Creating a new instance of http.ServeMux using the http.NewServeMux function
    mux := http.NewServeMux()

    // Register the custom 404 handler for all paths
    mux.HandleFunc("/", notFoundHandler)

    // Creating and configuring an HTTP server with the ServeMux instance as the handler
    server := &http.Server{
        Addr:    "localhost:8080",
        Handler: mux,
    }

    // Starting the server
    server.ListenAndServe()
}
