package models

import (
	"errors"
)

// use ErrNoRecord instead of sql.ErrNoRows to  encapsulate the model
// completely, so that our application isn’t concerned with the
// underlying datastore or reliant on datastore-specific errors for its
// behavior.
var ErrNoRecord = errors.New("models: no matching record found")
