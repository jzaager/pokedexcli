package main

import (
	"github.com/jzaager/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}
