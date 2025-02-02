package main

import "os"

func exit(cfg *config, args ...string) error {
	println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
