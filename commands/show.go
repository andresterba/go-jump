package commands

import (
	"fmt"
)

// SearchForDatabaseHit search for a hit in the database
func SearchForDatabaseHit(path string) (string, error) {
	foundPath, err := db.FindEntry(path)
	if err != nil {
		return foundPath, err
	}

	return foundPath, nil
}

// ShowCurrentEntriesInDatabase prints all entries in the database
func ShowCurrentEntriesInDatabase() error {
	entries, err := db.GetAllEntriesAsString()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Printf("%s\n", entry)
	}

	return nil
}
