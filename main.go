package main

import (
	"time"

	"github.com/jzaager/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: client,
	}

	startRepl(cfg)
}
