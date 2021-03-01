package database

import (
	"testing"
)

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

		testEntry := NewEntry(tt.input.counter, tt.input.path, "")

		if testEntry.isForPath(tt.input.searchPath) != tt.expectedOutput {
			t.Errorf("Expected search result is is not '%t'. got='%t'", tt.expectedOutput, testEntry.isForPath(tt.input.searchPath))
		}
	}
}
