package main

// struct with name & description strings
// the callback function gives the possibility to call the key command easely
type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			name:        "next cities",
			description: "Displays next 20 cities",
			callback:    mapn,
		},
		"mapb": {
			name:        "previous cities",
			description: "Displays previous 20 cities",
			// callback:    mapb,
		},
	}

}
