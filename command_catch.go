package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Must provide a pokemon name")
	}

	name := args[0]

	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	fmt.Println()

	return nil
}
