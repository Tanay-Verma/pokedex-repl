package main

import "fmt"

func commandHelp(cfg *config) error {
	command := getCommands()
	helpStr := "\nWelcome to the Pokedex!\nUsage:\n\n"
	for k, v := range command {
		commandDesc := fmt.Sprintf("%s: %s\n", k, v.description)
		helpStr += commandDesc
	}
	fmt.Println(helpStr)
	return nil
}
