package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	total := len(cfg.caughtPokemon)

	if total == 0 {
		fmt.Println("You have not caught any pokemon yet!")
		fmt.Println()
		return nil
	}

	fmt.Printf("You have caught %d pokemon!\n", total)
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	fmt.Println()
	return nil
}
