package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Usage: inspect <pokemon>")
	}
	name := args[0]
	info, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("Pokemon not on pokedex!")
	}

	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)
	fmt.Println("Stats:")
	for _, stat := range info.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range info.Types {
		fmt.Printf("  -%s\n", v.Type.Name)
	}

	return nil
}
