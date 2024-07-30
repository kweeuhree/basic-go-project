// import main package
package main

// import logging and net/http
import (
	"log"
	"net/http"
)

// define main point of entry
func main() {
	// initialize servemux
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static"
	// directory. Note that the path given to the http.Dir function is relative
	// to the project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the
	// handler for all URL paths that start with "/static/".
	// For matching paths, we strip the "/static" prefix before the request
	// reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//define handlers
	mux.HandleFunc("/", home)                        //catch-all subtree path
	mux.HandleFunc("/snippet/view", snippetView)     // fixed path
	mux.HandleFunc("/snippet/create", snippetCreate) // fixed path

	// log to console
	log.Print("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	// in case of errors log and exit
	log.Fatal(err)
}
