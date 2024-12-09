package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Tanay-Verma/pokedex-repl/internal/pokeapi"
)

type config struct {
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	pokeapiClient        pokeapi.Client
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		cmd := scanner.Text()

		command := getCommands()
		v, ok := command[cmd]
		if !ok {
			fmt.Println("Command not found")
			continue
		}

		if err := v.callback(cfg); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
