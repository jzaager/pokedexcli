package cmd

import (
	"fmt"

	"github.com/jzaager/pokedexcli/config"
)

func Help(cfg *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}
