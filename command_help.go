package main

import "fmt"

func commandHelp() {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	availableCommands := getCommand()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}
