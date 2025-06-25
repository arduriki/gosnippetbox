package main

import (
	"log"
	"net/http"
)

// home handler which writes a byte slice containing a response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Initialize a new servemux, then register the home function
	// as a handler for the "/" URL pattern.
    // Careful: it'll act as a 'catch-all' of any URL path
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Start a new web server on localserver: 'host:port'
	err := http.ListenAndServe(":4000", mux)
	// If an error is returned log the error message and terminate the program
	log.Fatal(err)
}
