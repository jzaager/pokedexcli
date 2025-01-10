package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/jzaager/pokedexcli/cmd"
	"github.com/jzaager/pokedexcli/config"
	"github.com/jzaager/pokedexcli/internal/pokeapi"
)

func cleanup() {
	keyboard.Close()
	fmt.Println("\nTerminal Reset. Exiting...")
}

func startRepl() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: client,
	}

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
			keyboard.Close()
			cmd.Exit(cfg)
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
			if inputBuffer.Len() == 0 {
				continue
			}
			handleBackspace(&inputBuffer)
			idx = len(commandHistory) - 1
		default:
			inputBuffer.WriteRune(char)
			fmt.Print(string(char))
		}
	}
}

func updateCommandHistory(commands []string, buf *strings.Builder) []string {
	commands = append(commands, buf.String())

	buf.Reset()
	fmt.Print("Pokedex > ")
	return commands
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

func runCommand(cfg *config.Config, commandName string, args ...string) error {
	if commandName == "" {
		return fmt.Errorf("\nNo command provided")
	}

	command, exists := cmd.GetCommands()[commandName]
	if !exists {
		return fmt.Errorf("\nUnknown command")
	}
	err := command.Callback(cfg, args...)
	if err != nil {
		return err
	}
	return nil
}
