package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() {
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

		// even if user inputs multiples words, only the first one will be used (for the time)
		cmd, exists := commands()[words[0]]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Take the input string and transform all chars to lower and create a slice of words separated by blanks
func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
