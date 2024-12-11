package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasURL = locationsResp.Next
	cfg.prevLocationAreasURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreasURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasURL = locationResp.Next
	cfg.prevLocationAreasURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
