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
		return fmt.Errorf("Could not find pokemon %q\n", pokemonName)
	}

	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %s...", pokemonName)
	fmt.Println()

	//	1000 - 600 = 400 // 400 / 1000 = 0.40
	//	1000 - 50  = 950 // 950 / 1000 = 0.95
	baseXP := pokemonResp.BaseExperience
	catchRate := float32(1000-baseXP) / float32(1000)
	esacpeRate := 1 - rand.Float32()

	/*
		alt for catch chance:
		catchChance := rand.Intn(pokemonResp.BaseExperience)
		if catchChance > 40 {
			// print escaped
			// return nil
		}
	*/

	if catchRate < esacpeRate {
		fmt.Printf("%s escaped!\n\n", pokemonName)
		return nil
	}

	fmt.Printf("%s has been caught!\n\n", pokemonName)

	cfg.caughtPokemon[pokemonName] = pokemonResp
	return nil
}
