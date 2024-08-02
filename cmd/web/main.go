// import main package
package main

// import logging and net/http
import (
	"database/sql"
	"flag" // handle command-line flags and arguments
	"fmt"
	"log"
	"net/http"
	"os" // operating system-level operations: handle files, directories, env variables, etc

	// environment variables
	"github.com/joho/godotenv"

	// we need the driver’s init() function to run so that it can register itself with the
	// database/sql package. The trick to getting around this is to alias the package name
	// to the blank identifier. This is standard practice for most of Go’s SQL drivers
	_ "github.com/go-sql-driver/mysql" // with underscore
)

// Define an application struct to hold the application-wide dependencies for
// the web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// The responsibilities of our main() function are limited to:
// - Parsing the runtime configuration settings for the application;
// - Establishing the dependencies for the handlers; and
// - Running the HTTP server.

// define main point of entry
func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DSN string with loaded env variables
	DSNstring := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbUser, dbPassword, dbName)

	// define  new command-line flag for the mysql dsn string
	dsn := flag.String("dsn", DSNstring, "MySQL data source name")

	// define a new command-line flag with name 'addr', a default value of ":4000"
	// and short help text explaining what the flag controls. The value of the flag
	// will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// use flag.Parse() function to parse the command-line flag
	// This reads in the command-line flag value and assigns it to the
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

	// create connection pool, pass openDB() the dsn from the command-line flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// defer a call to db.Close() so that the connection pool is closed before
	// the main() function exits
	defer db.Close()

	// initialize a new instance of application struct, containig the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initialize a new http.Server struct. Set the Addr and Handler fields so
	// that the server uses the same network address and routes as before,
	// and set the ErrorLog field so that the server now uses the custom errorLog
	// logger in the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		// Call the new app.routes() method to get the servemux containing routes.
		Handler: app.routes(),
	}

	// The value returned form the flag.String() function is a pointer to the flag
	// value, not the value itself. The dereference of the pointer is needed before the usage.
	//(prefix it with *)
	// use log.Printf() to interpolate the address with the log message
	// -- will also call os.Exit(1) after writing the message,
	// -- causing the application to immediately exit.
	infoLog.Printf("Starting server on %s", *addr)
	// use assignment operator as the err variable is already declared above
	err = srv.ListenAndServe()
	// in case of errors log and exit
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given dsn
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
