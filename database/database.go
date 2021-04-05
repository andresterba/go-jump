package database

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
	"time"
)

const hoursInAYear = 8760

var (
	ErrNotFound             = errors.New("item could not be found in the store")
	ErrCouldNotPersist      = errors.New("could not persist database")
	ErrCouldNotOpenFile     = errors.New("could not open database file")
	ErrCouldNotWriteFile    = errors.New("could not write database file")
	ErrWhileReadingDatabase = errors.New("could not read database")
)

// Database represents to database file
type Database struct {
	Path      string
	EntryList entryList
}

// NewDatabase reads a new database from a given path.
func NewDatabase(pathToDatabaseFile string) *Database {
	db := Database{Path: pathToDatabaseFile, EntryList: entryList{}}
	db.readDatabase()

	return &db
}

func (database *Database) sortDatabase() {
	sort.Sort(database.EntryList)
}

func (database *Database) readDatabase() error {
	// Flush EntryList as we can only read or write the hole database.
	// This ensures that the entries are only read once and therefore exists only
	// once in the database.
	database.flushEntryList()

	file, err := os.Open(database.Path)
	if err != nil {
		return ErrCouldNotOpenFile
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splittedLine := strings.Split(line, " ")

		dbEntry := NewEntry(splittedLine[0], splittedLine[1], splittedLine[2])
		database.EntryList = append(database.EntryList, dbEntry)
	}

	if err := fileScanner.Err(); err != nil {
		return ErrWhileReadingDatabase
	}

	return nil
}

// Persist stores the database in the filesystem
func (database *Database) Persist() error {

	database.sortDatabase()

	_ = os.Remove(database.Path)

	file, err := os.OpenFile(database.Path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return ErrCouldNotOpenFile
	}

	for _, databaseEntry := range database.EntryList {
		_, err = file.Write([]byte(databaseEntry.getWritableFormat() + "\n"))
		if err != nil {
			return ErrCouldNotWriteFile
		}
	}

	return nil
}

// AddEntry adds a new entry to the database and persists the database.
// It checks if an entry already exsists and increases its counter.
// If not a new entry will be created.
func (database *Database) AddEntry(path string) error {
	for _, databaseEntry := range database.EntryList {
		if databaseEntry.isForPath(path) {
			databaseEntry.incrementPathCounter()

			return nil
		}
	}

	newDatabaseEntry := NewEntry("1", path, "")
	database.EntryList = append(database.EntryList, newDatabaseEntry)

	err := database.Persist()
	if err != nil {
		return ErrCouldNotPersist
	}

	return nil
}

// FindEntry searches all entries for a give path or subpath.
func (database *Database) FindEntry(path string) (string, error) {
	for _, entry := range database.EntryList {
		if entry.isForPath(path) {
			return entry.Path, nil
		}
		if entry.isForPathInLowerCase(path) {
			return entry.Path, nil
		}
		if entry.isForPartOfThePath(path) {
			return entry.Path, nil
		}
		if entry.isForPartInLowerCaseOfThePath(path) {
			return database.Path, nil
		}
	}

	return "", ErrNotFound
}

// GetAllEntriesAsString gets all entries currently in the database as string.
func (database *Database) GetAllEntriesAsString() ([]string, error) {
	var currentDatabaseEntriesAsString []string

	for _, entry := range database.EntryList {
		currentDatabaseEntriesAsString = append(currentDatabaseEntriesAsString, entry.getWritableFormat())
	}

	return currentDatabaseEntriesAsString, nil
}

func (database *Database) flushEntryList() {
	database.EntryList = []*Entry{}
}

// PruneEntriesOlderOneYear prunes all images older than 1 year.
func (database *Database) PruneEntriesOlderOneYear() error {
	var toKeepEntries = entryList{}

	currentTime := time.Now().UTC()

	for _, entry := range database.EntryList {
		timeDifferenceInHours := currentTime.Sub(entry.LastVisit).Hours()

		if timeDifferenceInHours >= hoursInAYear {
			continue
		}

		toKeepEntries = append(toKeepEntries, entry)
	}

	database.EntryList = toKeepEntries

	err := database.Persist()
	if err != nil {
		return err
	}

	return nil
}
