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
