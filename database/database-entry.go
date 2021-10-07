package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	ErrCouldNotParseDate = errors.New("could not parse date")
)

// Entry represents a single record of a hit path in the database.
type Entry struct {
	Counter   int
	Path      string
	LastVisit time.Time
}

// NewEntry creates a new entry.
func NewEntry(counter string, path string, lastVisit string) *Entry {
	counterAsInt, err := strconv.Atoi(counter)
	if err != nil {
		fmt.Println(err)
	}

	if lastVisit == "" {
		return &Entry{Counter: counterAsInt, Path: path, LastVisit: time.Now()}
	}

	lastVisitParsed, err := time.Parse(time.RFC3339, lastVisit)
	if err != nil {
		log.Fatal(ErrCouldNotParseDate)
	}

	return &Entry{Counter: counterAsInt, Path: path, LastVisit: lastVisitParsed}
}

func (entry *Entry) incrementPathCounter() {
	entry.Counter++
	entry.LastVisit = time.Now()
}

func (entry Entry) getWritableFormat() string {
	return fmt.Sprintf("%d %s %s", entry.Counter, entry.Path, entry.LastVisit.Format(time.RFC3339))
}

func (entry Entry) isForPath(path string) bool {
	return entry.Path == path
}

func (entry Entry) isForPathInLowerCase(path string) bool {
	return entry.Path == strings.ToLower(path)
}

func (entry Entry) isForPartOfThePath(path string) bool {
	splittedPath := strings.Split(entry.Path, "/")
	numberOfFoldersInPath := len(splittedPath)

	for i := numberOfFoldersInPath - 1; i >= 0; i-- {
		if splittedPath[i] == path {
			return true
		}
	}

	return false
}

func (entry Entry) isForPartInLowerCaseOfThePath(path string) bool {
	splittedPath := strings.Split(entry.Path, "/")
	numberOfFoldersInPath := len(splittedPath)

	for i := numberOfFoldersInPath - 1; i >= 0; i-- {
		if strings.ToLower(splittedPath[i]) == path {
			return true
		}
	}

	return false
}
