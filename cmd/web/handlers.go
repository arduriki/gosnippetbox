package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Add home handler which writes a byte slice containing a response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Create a response header map:
	// the first parameter is the header name,
	// the second parameter is the header value.
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function to see the created snippets.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request
	// and try to convert it to an integer.
	id, err := strconv.Atoi(r.PathValue("id"))
	// If it can't be converted to an integer, or the value is less than 1,
	// return 404 page not found.
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
    
	// A message with the id value as the HTTP response.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

// Add a snippetCreatePost handler function.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Send a status created before the w.Write
	w.WriteHeader(http.StatusCreated)

	// Write the response body.
	w.Write([]byte("Save a new snippet..."))
}