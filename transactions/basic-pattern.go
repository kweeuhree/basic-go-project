type ExampleModel struct {
	DB *sql.DB
}

func (m *ExampleModel) ExampleTransaction() error {
	// Calling the Begin() method on the connection pool creates a new sql.Tx
	// object, which represents the in-progress database transaction.
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	// Defer a call to tx.Rollback() to ensure it is always called before the
	// function returns. If the transaction succeeds it will be already be
	// committed by the time tx.Rollback() is called, making tx.Rollback() a
	// no-op. Otherwise, in the event of an error, tx.Rollback() will
	// rollback the changes before the function returns.
	defer tx.Rollback()

	// Call Exec() on the transaction, passing in your statement and any
	// parameters. It's important to notice that tx.Exec() is called on the
	// transaction object just created, NOT the connection pool. Although
	// we're using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
	// exactly the same way.
	_, err = tx.Exec("INSERT INTO ...")
	if err != nil {
		return err
	}

	// Carry out another transaction in exactly the same way.
	_, err = tx.Exec("UPDATE ...")
	if err != nil {
		return err
	}

	// If there are no errors, the statements in the transaction can be
	// committed to the database with the tx.Commit() method.
	err = tx.Commit()

	return err
}

//You must always call either Rollback() or Commit()
// before your function returns. If you don’t the connection will
// stay open and not be returned to the connection pool. This can
// lead to hitting your maximum connection limit/running out of
// resources. The simplest way to avoid this is to use
// defer tx.Rollback() like we are in the example above.
// Transactions are also super-useful if you want to execute multiple
// SQL statements as a single atomic action. So long as you use the
// tx.Rollback() method in the event of any errors, the transaction
// ensures that either:
// All statements are executed successfully; or
// No statements are executed and the database remains
// unchanged.