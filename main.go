package main

import (
	"log"
	"net/http"
)

// Add home handler which writes a byte slice containing a response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function to see the created snippets.
func snippetView(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a specific snippet..."))
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
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)
    
	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Start a new web server on localserver: 'host:port'
	err := http.ListenAndServe(":4000", mux)
	// If an error is returned log the error message and terminate the program
	log.Fatal(err)
}
