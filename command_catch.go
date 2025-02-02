package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Usage: catch <pokemon>\n")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return errors.New("Pokemon doesn't exists\n")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	rate := rand.Intn(pokemon.BaseExperience)
	// fmt.Println("rate: ", rate)

	success := int(float64(pokemon.BaseExperience) * 0.50)
	// fmt.Println("success: ", success)

	if rate >= success {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Println("Pokemons available in your pokedex!")

	// fmt.Println(pokemon)

	return nil
}
