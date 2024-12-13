package main

type cliCommand struct {
	callback    func(cfg *config, args ...string) error
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
		"explore": {
			name:        "explore",
			description: "Explores the given area and returns the Pokemon found there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch the pokemon.",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Print the Pokemon's stat",
			callback:    commandInspect,
		},
	}
	return command
}
