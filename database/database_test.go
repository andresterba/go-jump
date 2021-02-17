package database

import (
	"os/exec"
	"testing"
)

const (
	testDatabasePath = "test.db"
)

func removeTestDatabaseFile() error {
	runCommand := exec.Command("rm", testDatabasePath)
	err := runCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func TestDatabaseContainsCorrectNumberOfEntries(t *testing.T) {
	type entryInput struct {
		paths []string
	}

	type expectedOutput struct {
		numberOfEntries int
		path            string
	}

	tests := []struct {
		input          entryInput
		expectedOutput expectedOutput
	}{
		{entryInput{[]string{"test"}}, expectedOutput{1, ""}},
		{entryInput{[]string{"test1", "test2"}}, expectedOutput{2, ""}},
		{entryInput{[]string{"test1", "test2", "test3", "test4", "test5"}}, expectedOutput{5, ""}},
	}

	for _, tt := range tests {

		testDatabase := NewDatabase(testDatabasePath)

		for _, inputPath := range tt.input.paths {
			testDatabase.AddEntry(inputPath)
		}

		entries := testDatabase.EntryList

		if entries.Len() != tt.expectedOutput.numberOfEntries {
			t.Errorf("Number of entries is not %d. got=%d", tt.expectedOutput.numberOfEntries, entries.Len())
		}

		err := removeTestDatabaseFile()
		if err != nil {
			t.Fatalf("Could not delete test database. got=%s", err)
		}
	}
}