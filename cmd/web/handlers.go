// import main package
package main

// import formatting functions
import (
	"fmt"
	"net/http"
	"strconv"
)

// define home function handler
func home(w http.ResponseWriter, r *http.Request) {
	// check if the page exists
	if r.URL.Path != ("/") {
		http.NotFound(w, r)
		return
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
