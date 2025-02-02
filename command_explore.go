package main

import (
	"errors"
	"fmt"
)

func exploreArea(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Usage: explore <location>\n")
	}

	name := args[0]
	zoneResp, err := cfg.pokeapiClient.LocationPokemons(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args)
	fmt.Println("Found Pokemon:")
	for _, pok := range zoneResp.PokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}
	return nil
}
