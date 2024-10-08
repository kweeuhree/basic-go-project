package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
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
	//route declarations use application struct's methods as the handler functions
	mux.HandleFunc("/", app.home)                        //catch-all subtree path
	mux.HandleFunc("/snippet/view", app.snippetView)     // fixed path
	mux.HandleFunc("/snippet/create", app.snippetCreate) // fixed path

	return mux
}
