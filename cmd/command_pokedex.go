package cmd

import (
	"fmt"

	"github.com/jzaager/pokedexcli/config"
)

func Pokedex(cfg *config.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	total := len(cfg.CaughtPokemon)

	if total == 0 {
		fmt.Println("You have not caught any pokemon yet!")
		fmt.Println()
		return nil
	}

	fmt.Printf("You have caught %d pokemon!\n", total)
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	fmt.Println()
	return nil
}
