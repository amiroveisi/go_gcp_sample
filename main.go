package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment variable, default to 8080
	// This allows Kubernetes to configure the port via environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for local development and container
	}

	// Main application endpoint - responds to all requests to "/"
	// This is a simple handler that returns a greeting message with version info
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go app! Version: 1.0.0\n")
	})

	// Health check endpoint - required for Kubernetes liveness and readiness probes
	// Kubernetes will periodically call this endpoint to check if the pod is healthy
	// Returns HTTP 200 OK status to indicate the application is running properly
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	// Log server startup for debugging and monitoring
	log.Printf("Server starting on port %s", port)

	// Start HTTP server and listen on the specified port
	// ListenAndServe blocks and runs the server until it encounters an error
	// log.Fatal will terminate the program if the server fails to start
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
