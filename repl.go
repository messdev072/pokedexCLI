package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommand()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
