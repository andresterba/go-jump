package commands

import (
	"fmt"
)

func GetHit(path string) (string, error) {
	foundPath, err := db.FindEntry(path)
	if err != nil {
		return foundPath, err
	}

	return foundPath, nil
}

func Show() error {

	entries, err := db.GetAllEntriesAsString()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Printf("%s\n", entry)
	}

	return nil
}
