package main

import "fmt"

// Print help page
// Iterate over commands map where the key is "help"
// Print informations to stdout
func help() error {
	fmt.Println("\nWelcome to Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range commands() {
		fmt.Printf("%s: %s\n", v.name, v.description)

	}
	fmt.Println()
	return nil
}
