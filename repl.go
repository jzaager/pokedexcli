package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jzaager/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type CommandRegistry map[string]cliCommand

var supportedCommands CommandRegistry

func getCommands() CommandRegistry {
	return CommandRegistry{
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
			description: "Display next page of locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page of locations",
			callback:    commandMapB,
		},
	}
}
