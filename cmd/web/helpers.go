package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
// -- use the debug.Stack() function to get a stack trace for the current goroutine and append it to the
// -- log message. Being able to see the execution path of the
// -- application via the stack trace can be helpful when you’re trying to debug errors.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// report the file name and line number one step back in the stack trace
	// to have a clearer idea of where the error actually originated from
	// set frame depth to 2
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400
// "Bad Request" when there's a problem with the request that the user sent.
// -- use the http.StatusText() function to automatically generate a human-friendly text
// representation of a given HTTP status code. For example,
// http.StatusText(400) will return the string "Bad Request".
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found
// response to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
