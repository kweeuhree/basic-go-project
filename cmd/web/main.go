// import main package
package main

// import logging and net/http
import (
	"flag" // handle command-line glags and arguments
	"log"
	"net/http"
	"os" // operating system-level operations: handle files, directories, env variables, etc
)

// define main point of entry
func main() {

	// define a new command-line flag with name 'addr', a default value of ":4000"
	// and short help text explaining what the flag controls. The value of the flag
	// will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// use flag.Parse() function to parse the command-line flag
	// This reads in the command-line glaf value and assigns it to the
	// addr variable. You need to call this _before_ you use the addr variable,
	// otherwise it will always contain the default value of ":4000"
	// application will be terminated in case of any errors
	flag.Parse()

	// Use log.New() to create a logger for writing information messages.
	// Parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate
	// what additional information to include (local date and time).
	// The flags are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use
	// 	stderr as the destination and use the log.Lshortfile flag to include the
	// relevant file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize a new servemux
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

	// Initialize a new http.Server struct. Set the Addr and Handler fields so
	// that the server uses the same network address and routes as before,
	// and set the ErrorLog field so that the server now uses the custom errorLog
	// logger in the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// The value returned form the flag.String() function is a pointer to the flag
	// value, not the value itself. The dereference of the pointer is needed before the usage.
	//(prefix it with *)
	// use log.Printf() to interpolate the address with the log message
	// -- will also call os.Exit(1) after writing the message,
	// -- causing the application to immediately exit.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	// in case of errors log and exit
	errorLog.Fatal(err)
}
