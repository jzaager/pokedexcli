package main

import (
	"time"

	"github.com/jzaager/pokedexcli/internal/pokeapi"
	"github.com/jzaager/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(5*time.Second, cache)
	cfg := &config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}
