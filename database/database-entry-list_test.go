package database

import (
	"sort"
	"testing"
)

func TestDatabaseEntryListSort(t *testing.T) {
	type entryInput struct {
		counter string
		path    string
	}

	tests := []struct {
		input          []entryInput
		expectedOutput entryList
	}{
		{[]entryInput{
			{"1", "/home/test1"},
			{"2", "/home/test2"},
			{"3", "/home/test3"},
		},
			entryList{
				&Entry{Counter: 3, Path: "/home/test3"},
				&Entry{Counter: 2, Path: "/home/test2"},
				&Entry{Counter: 1, Path: "/home/test1"},
			},
		},
	}

	for _, tt := range tests {
		var entryListTest entryList

		for _, inputEntry := range tt.input {
			testEntry := NewEntry(inputEntry.counter, inputEntry.path)
			entryListTest = append(entryListTest, testEntry)
		}

		sort.Sort(entryListTest)

		for i := 0; i < len(entryListTest); i++ {
			if entryListTest[i].Counter != tt.expectedOutput[i].Counter && entryListTest[i].Path != tt.expectedOutput[i].Path {
				t.Errorf("Entry at position %d is not %s. got=%s",
					i,
					tt.expectedOutput[i].getWritableFormat(),
					entryListTest[i].getWritableFormat(),
				)
			}
		}
	}
}
