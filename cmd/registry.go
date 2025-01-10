package cmd

import (
	"github.com/jzaager/pokedexcli/config"
)

type CliCommand struct {
	name        string
	description string
	Callback    func(cfg *config.Config, args ...string) error
}

type CommandRegistry map[string]CliCommand

var supportedCommands CommandRegistry

func GetCommands() CommandRegistry {
	return CommandRegistry{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"map": {
			name:        "map",
			description: "Displays next page of locations",
			Callback:    MapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of locations",
			Callback:    MapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays a list of pokemon at a given location",
			Callback:    Explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			Callback:    Catch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught pokemon",
			Callback:    Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of all your caught pokemon",
			Callback:    Pokedex,
		},
	}
}
