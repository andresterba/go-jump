package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andresterba/go-jump/commands"
)

func printHelp() {
	fmt.Printf(`
go-jump [command]

commands:
    * folder - search for best match for "folder"
    * add    - add a folder to the database
    * show   - show current database entries
    * prune  - prune entries older than 1 year
`)
}

const (
	databaseFileName = ".go-jump.db"
)

var (
	errInputNotValid = errors.New("input is not valid")
)

// isInputValid checks if a given input contains at least
// a single slash (which makes it a valid path) AND is not $HOME.
func isInputValid(input string, userHomeDir string) bool {
	if !strings.Contains(input, "/") {
		return false
	}

	if input == userHomeDir {
		return false
	}

	return true
}

func checkErrorAndLogAndFail(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) <= 1 {
		printHelp()
		os.Exit(1)
	}

	command := args[1]
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	commands.RegisterDatabase(userHomeDir + "/" + databaseFileName)

	switch command {
	case "add":
		if len(args) <= 2 {
			printHelp()
			os.Exit(1)
		}

		path := args[2]
		if isInputValid(path, userHomeDir) {
			err := commands.Add(path)
			if err != nil {
				return err
			}

			break
		}

	case "show":
		err = commands.ShowCurrentEntriesInDatabase()
		if err != nil {
			return err
		}
		break

	case "prune":
		err = commands.Prune()
		if err != nil {
			return err
		}
		break

	default:
		foundPath, err := commands.SearchForDatabaseHit(command)
		checkErrorAndLogAndFail(err)

		fmt.Printf("%s\n", foundPath)

	}

	return nil
}

func main() {
	err := run(os.Args)
	checkErrorAndLogAndFail(err)
}
