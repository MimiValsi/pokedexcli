package main

import (
	"strings"
)

// struct with name & description strings
// the callback function gives the possibility to call the key command easely
type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

// commands function which returns a map of string key with cliCommand struct values
// Faster way of returning key-values information
func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 cities",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 cities",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "list all pokemon in the area",
			callback:    exploreArea,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Tries to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Show information about caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all pokemons caught",
			callback:    commandPokedex,
		},
	}

}

// Take the input string and transform all chars to lower and create a slice of words separated by blanks
func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
