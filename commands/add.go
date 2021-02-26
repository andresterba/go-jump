package commands

// Add adds a path to the database and persists the database.
func Add(pathToAdd string) error {
	err := db.AddEntry(pathToAdd)
	if err != nil {
		return err
	}

	err = db.Persist()
	if err != nil {
		return err
	}

	return nil
}
