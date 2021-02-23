package commands

import "github.com/andresterba/go-jump/database"

var db *database.Database

// RegisterDatabase creates a new database from a given path
func RegisterDatabase(path string) {
	db = database.NewDatabase(path)
}
