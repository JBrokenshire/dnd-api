package main

import (
	"dnd-api/db/migrations/process"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

var confirmationMessage = "\nAre you sure you want to update the database? Y or N:  "

func init() {
	flag.Bool("confirm", false, "Should we ask for confirmation?")
}

func main() {
	confirmRequired := os.Args[1] == "--confirm"
	if !confirmRequired {
		process.Run()
		return
	}

	if askForConfirmation() {
		log.Print()
		process.Run()
	}
}

func askForConfirmation() bool {
	fmt.Print(confirmationMessage)

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}

	if slices.Contains(okayResponses, response) {
		return true
	} else if slices.Contains(nokayResponses, response) {
		return false
	} else {
		fmt.Print(confirmationMessage)
		return askForConfirmation()
	}
}
