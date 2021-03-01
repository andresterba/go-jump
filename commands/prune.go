package commands

// Prune entries older than 1 year.
func Prune() error {
	err := db.PruneEntriesOlderOneYear()
	if err != nil {
		return err
	}

	return nil
}
