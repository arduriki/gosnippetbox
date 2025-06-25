package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Add home handler which writes a byte slice containing a response body.
func home(w http.ResponseWriter, r *http.Request) {
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
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// Initialize a new servemux, then register the home function
	// as a handler for the "/" URL pattern.
	mux := http.NewServeMux()
    // {$} to prevent subtree path, it helps to match a single slash with
    // nothing else.
	mux.HandleFunc("/{$}", home)
    // Register view and create with their corresponding route patterns.
    mux.HandleFunc("/snippet/view/{id}", snippetView) // {id} is a wildcard
    mux.HandleFunc("/snippet/create", snippetCreate)
    
	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Start a new web server on localserver: 'host:port'
	err := http.ListenAndServe(":4000", mux)
	// If an error is returned log the error message and terminate the program
	log.Fatal(err)
}
