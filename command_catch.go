package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("Could not find pokemon %s\n", pokemonName)
	}

	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	fmt.Println()

	// TODO: use pokemonResp
	baseXP := pokemonResp.BaseExperience
	//		     1000 - 600 = 400 // 400 / 1000 = 0.40
	//			 1000 - 50  = 950 // 950 / 1000 = 0.95
	catchRate := float32(1000-baseXP) / float32(1000)
	esacpeRate := 1 - rand.Float32()

	fmt.Printf("Catch rate: %.2f\n", catchRate)
	fmt.Printf("Escape rate: %.2f\n", esacpeRate)

	if catchRate > esacpeRate {
		fmt.Printf("%s has been caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	fmt.Println()
	return nil
}
