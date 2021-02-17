package commands

import "github.com/andresterba/go-jump/database"

type Command interface {
	Execute() error
}

var db *database.Database

func registerDatabase() {
	db = database.NewDatabase("/home/aps/example.db")
}

func init() {
	registerDatabase()
}
