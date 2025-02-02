package main

import "fmt"

// Print help page
// Iterate over commands map where the key is "help"
// Print informations to stdout
func help(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	fmt.Println()
	return nil
}
