// import main package
package main

// import formatting functions
import (
	"fmt"
	"net/http"
	"strconv"

	// provides a family of functions for safely parsing and rendering HTML
	// templates. We can use the functions in this package to parse the
	// template file and then execute the template.
	"html/template"
)

// define home function handler
// signature of the home handler is defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if the page exists
	if r.URL.Path != ("/") {
		app.notFound(w) // use notFound() helper
		return
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	// Use the template.ParseFiles() function to read the template file into
	// a template set(ts). If there's an error, log the detailed error message
	// and use the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// -- files can be an absolute path, a relative path,
	// -- or multiple file paths passed as variadic arguments
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// because the home hadnler function is a method against application
		// it can access its fields, including the error logger
		app.serverError(w, err) // use serverError() helper
		return
	}
	// use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in.
	// -- use the ExecuteTemplate() method to respond using the content
	// -- of the base template (which in turn invokes title and main templates).
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		//error logger form the application struct
		app.notFound(w) // use notFound() helper
		return
	}

	w.Write([]byte("hello byte"))
}

// define view snippet function
// signature of the handler is defined to be a method against the *application struct
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// get id from request URL query
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// in case of invalid id or error return a 404 NotFound
	if err != nil || id < 1 {
		app.notFound(w) // use notFound() helper
		// return to prevent further code execution
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// define create a snippet function
// signature is defined to be a method against the *application struct
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		// call a http.Error shortcut
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		// use clientError helper instead of a http.Error shortcut
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet"))
}
