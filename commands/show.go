package commands

import "fmt"

func GetHit(path string) (string, error) {
	foundPath, err := db.FindEntry(path)
	if err != nil {
		return foundPath, err
	}

	return foundPath, nil
}

func Show() {
	fmt.Println("Not implemented yet!")
}
