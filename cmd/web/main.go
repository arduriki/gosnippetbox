package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new servemux, then register the home function
	// as a handler for the "/" URL pattern.
	mux := http.NewServeMux()
	// {$} to prevent subtree path, it helps to match a single slash with
	// nothing else.
	mux.HandleFunc("GET /{$}", home) // GET to restrict the HTTP method.
	// Register view and create with their corresponding route patterns.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // {id} is a wildcard.
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// A POST route request only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Start a new web server on localserver: 'host:port'
	err := http.ListenAndServe(":4000", mux)
	// If an error is returned log the error message and terminate the program
	log.Fatal(err)
}
