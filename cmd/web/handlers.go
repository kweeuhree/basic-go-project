// import main package
package main

// import formatting functions
import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	// provides a family of functions for safely parsing and rendering HTML
	// templates. We can use the functions in this package to parse the
	// template file and then execute the template.
	"html/template"
)

// define home function handler
func home(w http.ResponseWriter, r *http.Request) {
	// check if the page exists
	if r.URL.Path != ("/") {
		http.NotFound(w, r)
		return
	}
	// Use the template.ParseFiles() function to read the template file into
	// a template set(ts). If there's an error, log the detailed error message
	// and use the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html") // can also be an absolute path
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	w.Write([]byte("hello byte"))
}

// define view snippet function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// get id from request URL query
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// in case of invalid id or error return a 404 NotFound
	if err != nil || id < 1 {
		http.NotFound(w, r)
		// return to prevent further code execution
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// define create a snippet function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		// call a http.Error shortcut
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet"))
}
