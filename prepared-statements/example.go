// We need somewhere to store the prepared statement for the lifetime of our
// web application. A neat way is to embed in the model alongside the
// connection pool.
type ExampleModel struct {
	DB *sql.DB
	InsertStmt *sql.Stmt
}

// Create a constructor for the model, in which we set up the prepared
// statement.
func NewExampleModel(db *sql.DB) (*ExampleModel, error) {
	// Use the Prepare method to create a new prepared statement for the
	// current connection pool. This returns a sql.Stmt object which
	// represents the prepared statement.
	insertStmt, err := db.Prepare("INSERT INTO ...")
	
	if err != nil {
		return nil, err
	}

	// Store it in our ExampleModel object, alongside the connection pool.
	return &ExampleModel{db, insertStmt}, nil
}

// Any methods implemented against the ExampleModel object will have access
// to the prepared statement.
func (m *ExampleModel) Insert(args...) error {
	// Notice how we call Exec directly against the prepared statement,
	// rather than against the connection pool? Prepared statements also support the
	// Query and QueryRow methods.
	_, err := m.InsertStmt.Exec(args...)
		return err
	}

// In the web application's main function we will need to initialize a new
// ExampleModel struct using the constructor function.	
func main() {
	db, err := sql.Open(...)
	
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	
	// Create a new ExampleModel object, which includes the prepared statement.
	exampleModel, err := NewExampleModel(db)
	if err != nil {
		errorLog.Fatal(err)
	}
	
	// Defer a call to Close() on the prepared statement to ensure that it is
	// properly closed before our main function terminates.
	defer exampleModel.InsertStmt.Close()
}
