package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		cmdInput := strings.Fields(scanner.Text())
		if len(cmdInput) == 0 {
			continue
		}

		cmd := cmdInput[0]
		args := []string{}
		if len(cmdInput) > 1 {
			args = cmdInput[1:]
		}

		command := getCommands()
		v, ok := command[cmd]
		if !ok {
			fmt.Println("Command not found")
			continue
		}

		if err := v.callback(cfg, args...); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
