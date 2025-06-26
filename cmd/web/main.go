package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// New command-line flag and some short help text of what is doing.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse the command-line flag.
	flag.Parse()

	// Initialize a new servemux, then register the home function
	// as a handler for the "/" URL pattern.
	mux := http.NewServeMux()

	// File server to serve the static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server and strip "/static" prefix before
	// the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// {$} to prevent subtree path, it helps to match a single slash with
	// nothing else.
	mux.HandleFunc("GET /{$}", home) // GET to restrict the HTTP method.
	// Register view and create with their corresponding route patterns.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // {id} is a wildcard.
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// A POST route request only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say that the server is starting.
	log.Printf("starting server on %s", *addr)

	// Start a new web server on localserver: 'host:port'
	err := http.ListenAndServe(*addr, mux)
	// If an error is returned log the error message and terminate the program
	log.Fatal(err)
}
