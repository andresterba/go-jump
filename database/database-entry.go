package database

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry represents a single record of a hit path in the database.
type Entry struct {
	Counter int
	Path    string
}

// NewEntry creates a new entry.
func NewEntry(counter string, path string) *Entry {
	counterAsInt, err := strconv.Atoi(counter)
	if err != nil {
		fmt.Println(err)
	}

	return &Entry{Counter: counterAsInt, Path: path}
}

func (entry Entry) getWritableFormat() string {
	return fmt.Sprintf("%d %s", entry.Counter, entry.Path)
}

func (entry Entry) isForPath(path string) bool {
	if entry.Path == path {
		return true
	}

	return false
}

func (entry Entry) isForPathInLowerCase(path string) bool {
	if entry.Path == strings.ToLower(path) {
		return true
	}

	return false
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
