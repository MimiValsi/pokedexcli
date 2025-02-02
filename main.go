package main

import (
	"time"

	"github.com/MimiValsi/pokedexcli/internal/pokeapi"
)

func main() {

	// cache := pokeapi.NewCache(5 * time.Minute)
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	start(cfg)
}
