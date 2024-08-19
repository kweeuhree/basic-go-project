package main

import "kweeuhree.snippetbox/internal/models"

// wrap your dynamic data in a struct which acts like
// a single ‘holding structure’ for your data.

// Define a templateData type to act as the holding structure for any dynamic data that we want
// to pass to our HTML templates. At the moment it only contains one field

type templateData struct {
	Snippet *models.Snippet
}
