// import main package
package main

// import logging and net/http
import (
	"flag" // handle command-line glags and arguments
	"log"
	"net/http"
)

// define main point of entry
func main() {

	// define a new command-line flag with name 'addr', a adefault value of ":4000"
	// and short help text explaining what the flag controls. The value of the flag
	// will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network addres")

	// use flag.Parse() function to parse the command-line flag
	// This reads in the command-line glaf value and assigns it to the
	// addr variable. You need to call this _before_ you use the addr variable,
	// otherwise it will always contain the default value of ":4000"
	// application will be terminated in case of any errors
	flag.Parse()

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

	// The value returned form the flag.String() function is a pointer to the flag
	// value, not the value itself. The dereference of the pointer is needed before the usage.
	//(prefix it with *)
	// use log.Printf() to interpolate the address with the log message
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	// in case of errors log and exit
	log.Fatal(err)
}
