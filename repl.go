package main

import (
	"fmt"
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
	commandHistory := []string{}
	idx := 0
	prevCommand := ""

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
		}

		switch key {
		case keyboard.KeyEsc:
			commandExit(cfg)
			return
		case keyboard.KeyArrowUp:
			prevCommand, idx = getPrevCommand(commandHistory, idx)
			printPrevCommand(&inputBuffer, prevCommand)
		case keyboard.KeySpace:
			inputBuffer.WriteRune(' ')
			fmt.Print(" ")
		case keyboard.KeyEnter:
			cmd, args, err := handleEnter(&inputBuffer)
			if err != nil {
				inputBuffer.Reset()
				fmt.Print("Pokedex > ")
				continue
			}
			err = runCommand(cfg, cmd, args...)
			if err != nil {
				fmt.Println(err)
				commandHistory = updateCommandHistory(commandHistory, &inputBuffer)
				idx = len(commandHistory) - 1
				continue
			}
			commandHistory = updateCommandHistory(commandHistory, &inputBuffer)
			idx = len(commandHistory) - 1
		case 127: // 'backspace'
			if inputBuffer.Len() > 0 {
				handleBackspace(&inputBuffer)
			}
			idx = len(commandHistory) - 1
		default:
			inputBuffer.WriteRune(char)
			fmt.Print(string(char))
		}
	}
}

func handleBackspace(buf *strings.Builder) {
	text := buf.String()
	buf.Reset()
	buf.WriteString(text[:len(text)-1])

	// Move cursor back, overwrite last char, and move back again
	fmt.Print("\b \b")
}

func updateCommandHistory(commands []string, buf *strings.Builder) []string {
	commands = append(commands, buf.String())

	buf.Reset()
	fmt.Print("Pokedex > ")
	return commands
}

func handleEnter(buf *strings.Builder) (string, []string, error) {
	fmt.Println()
	text := buf.String()
	if text == "" {
		return "", nil, fmt.Errorf("No text provided")
	}

	cmd, args := splitCommandArgs(text)
	return cmd, args, nil
}

func printPrevCommand(buf *strings.Builder, command string) {
	buf.Reset()
	buf.WriteString(command)
	fmt.Printf("\r\033[KPokedex > %s", command)
}

func getPrevCommand(allCommands []string, i int) (string, int) {
	if len(allCommands) == 0 {
		return "", i
	}
	if i < 0 {
		return allCommands[0], 0
	}
	return allCommands[i], i - 1
}

func runCommand(cfg *config, cmd string, args ...string) error {
	if cmd == "" {
		return fmt.Errorf("\nNo command provided")
	}
	command, exists := getCommands()[cmd]
	if !exists {
		return fmt.Errorf("\nUnknown command")
	}
	err := command.callback(cfg, args...)
	if err != nil {
		return err
	}
	return nil
}

func splitCommandArgs(text string) (command string, args []string) {
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
