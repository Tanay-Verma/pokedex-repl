package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide Pokemon name")
	}
	pokemonName := args[0]
	pokeInfo, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokeInfo.Name)
	fmt.Println("Height:", pokeInfo.Height)
	fmt.Println("Weight:", pokeInfo.Weight)

	fmt.Println("Stats:")
	for _, v := range pokeInfo.Stats {
		fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Println("Types:")
	for _, v := range pokeInfo.Types {
		fmt.Println("  -", v.Type.Name)
	}
	return nil
}
