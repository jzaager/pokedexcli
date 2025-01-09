package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/jzaager/pokedexcli/internal/pokeapi"
)

type config struct {
	caughtPokemon   map[string]pokeapi.Pokemon
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func cleanup() {
	keyboard.Close()
	fmt.Println("\nTerminal Reset. Exiting...")
}

func startRepl(cfg *config) {
	if err := keyboard.Open(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer cleanup()

	fmt.Print("Pokedex > ")
	var inputBuffer strings.Builder

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
		}

		switch key {
		case keyboard.KeyEsc:
			fmt.Println("Exiting...")
			cmd := exec.Command("stty", "sane")
			cmd.Stdin = os.Stdin
			cmd.Run()
			os.Exit(0)
			return
		case keyboard.KeyArrowUp:
			fmt.Print("Up ARROW PRESS")
		case keyboard.KeySpace:
			inputBuffer.WriteRune(' ')
			fmt.Print(" ")
		case keyboard.KeyEnter:
			text := inputBuffer.String()
			cmd, args := processText(text)
			err := runCommand(cfg, cmd, args...)
			if err != nil {
				fmt.Println(err)
			}
			inputBuffer.Reset()
			fmt.Print("Pokedex > ")
			continue
		case 127: // 'backspace'
			if inputBuffer.Len() > 0 {
				text := inputBuffer.String()
				inputBuffer.Reset()
				inputBuffer.WriteString(text[:len(text)-1])

				// Move cursos back, overwrite last char, and move back again
				fmt.Print("\b \b")
			}
		default:
			inputBuffer.WriteRune(char)
			fmt.Print(string(char))
		}
	}
}

func runCommand(cfg *config, cmd string, args ...string) error {
	if cmd == "" {
		return fmt.Errorf("\nNo command provided")
	}
	command, exists := getCommands()[cmd]
	if exists {
		command.callback(cfg, args...)
		return nil
	}
	return fmt.Errorf("\nUnknown command")
}

func processText(text string) (command string, args []string) {
	words := cleanInput(text)
	if len(words) == 0 {
		return "", nil
	}

	command = words[0]
	args = []string{}
	if len(words) > 1 {
		args = words[1:]
	}

	return command, args
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
	callback    func(cfg *config, args ...string) error
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
			description: "Displays next page of locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays a list of pokemon at a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of all your caught pokemon",
			callback:    commandPokedex,
		},
	}
}
