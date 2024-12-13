package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide Pokemon name")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience) * 100 / pokemon.BaseExperience
	if chance < 75.0 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
