package database

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strings"
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
	file, err := os.Open(database.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splittedLine := strings.Split(line, " ")
		dbEntry := NewEntry(splittedLine[0], splittedLine[1])
		// fmt.Printf("%v", dbEntry)
		database.EntryList = append(database.EntryList, dbEntry)
	}

	if err := fileScanner.Err(); err != nil {
		return errors.New("error while reading database")
	}

	return nil
}

// Persist stores the database in the filesystem
func (database *Database) Persist() error {

	database.sortDatabase()

	err := os.Remove(database.Path)

	file, err := os.OpenFile(database.Path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	for _, databaseEntry := range database.EntryList {
		//fmt.Printf("%s\n", databaseEntry.getWritableFormat())

		_, err = file.Write([]byte(databaseEntry.getWritableFormat() + "\n"))
		if err != nil {
			log.Fatal(err)
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
			databaseEntry.Counter++

			return nil
		}
	}

	newDatabaseEntry := NewEntry("1", path)
	database.EntryList = append(database.EntryList, newDatabaseEntry)

	err := database.Persist()
	if err != nil {
		return err
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

	return "", errors.New("not found")
}
