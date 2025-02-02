package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MimiValsi/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func start(cfg *config) {
	// Create variable which will be a *Scanner while reading from STDIN
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// Waits for user input and scans for inputs tokens
		input.Scan()
		// commands() returns a map[string]cliCommand
		// While a function is a first-class "citizen" we can call the function while adding the key for the map.
		// Exemple:
		// Pokedex > help
		// "help" exists in our cliCommand struct, therefore we call the function specified in the struct to print help page
		words := cleanInput(input.Text())
		if len(words) == 0 {
			fmt.Println("Need input!")
			continue
		}

		// even if user inputs multiples words, only the first one will be used (for the time)
		cmd, exists := commands()[words[0]]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")
		continue
	}
}
