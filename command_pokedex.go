package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.pokedex {
		fmt.Println(" -", key)
	}
	return nil
}
