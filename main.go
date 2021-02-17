package main

import (
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
`)
}

// isInputValid checks if a given input contains at least
// a single slash (which makes it a valid path).
func isInputValid(input string) bool {
	return strings.Contains(input, "/")
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

	switch command {
	case "add":
		if len(args) <= 2 {
			printHelp()
			os.Exit(1)
		}

		path := args[2]
		if isInputValid(path) {
			err := commands.Add(path)
			checkErrorAndLogAndFail(err)

			break
		}

		fmt.Printf("Please provide a valid input!")
		break

	case "show":
		commands.Show()
		break

	default:
		foundPath, err := commands.GetHit(command)
		checkErrorAndLogAndFail(err)

		fmt.Printf("%s\n", foundPath)

	}

	return nil
}

func main() {
	run(os.Args)
}
