package main

type cliCommand struct {
	callback    func(cfg *config) error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	command := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map back",
			description: "Displays the names of previous 20 location areas in Pokemon world",
			callback:    commandMapb,
		},
	}
	return command
}
