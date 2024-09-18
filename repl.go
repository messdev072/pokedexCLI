package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")

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
		err := command.callback(cfg)
		if err != nil {
			fmt.Print(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "Display 20 locations in Pokemon World",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display 20 previous locations in Pokemon World",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
