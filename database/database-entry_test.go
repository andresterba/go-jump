package database

import (
	"testing"
)

func TestDatabaseEntryPrint(t *testing.T) {
	type entryInput struct {
		counter string
		path    string
	}

	tests := []struct {
		input          entryInput
		expectedOutput string
	}{
		{entryInput{"1", "/home/test"}, "1 /home/test"},
		{entryInput{"2", "/home/test/documents"}, "2 /home/test/documents"},
		{entryInput{"3", "/home/test/myfile"}, "3 /home/test/myfile"},
		{entryInput{"4", "/home"}, "4 /home"},
		{entryInput{"5", "/home/test/Downloads"}, "5 /home/test/Downloads"},
	}

	for _, tt := range tests {

		testEntry := NewEntry(tt.input.counter, tt.input.path)

		if testEntry.getWritableFormat() != tt.expectedOutput {
			t.Errorf("Writable output is not %s. got=%s", tt.expectedOutput, testEntry.getWritableFormat())
		}
	}
}

func TestDatabaseEntryIsForPath(t *testing.T) {
	type entryInput struct {
		counter    string
		path       string
		searchPath string
	}

	tests := []struct {
		input          entryInput
		expectedOutput bool
	}{
		{entryInput{"1", "/home/test", "/home/test"}, true},
		{entryInput{"1", "/home/test", "test"}, false},
	}

	for _, tt := range tests {

		testEntry := NewEntry(tt.input.counter, tt.input.path)

		if testEntry.isForPath(tt.input.searchPath) != tt.expectedOutput {
			t.Errorf("Expected search result is is not %t. got=%t", tt.expectedOutput, testEntry.isForPath(tt.input.searchPath))
		}
	}
}
