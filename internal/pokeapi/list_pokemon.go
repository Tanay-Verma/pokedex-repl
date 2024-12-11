package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreaPokemons(locationArea string) (LocationAreaPokemonResp, error) {
	url := baseURL + "/location-area/" + locationArea

	var locationAreaPokemonResp LocationAreaPokemonResp

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &locationAreaPokemonResp); err != nil {
			return locationAreaPokemonResp, err
		}
		return locationAreaPokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaPokemonResp, fmt.Errorf("Error in creating NewRequest: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaPokemonResp, fmt.Errorf("Error in calling the url: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaPokemonResp, fmt.Errorf("Error in loading the data to memory: %v", err)
	}

	if resp.StatusCode != 200 {
		return locationAreaPokemonResp, fmt.Errorf("%v in fetching the pokemons. Please enter a valid location-area", resp.StatusCode)
	}

	if err := json.Unmarshal(data, &locationAreaPokemonResp); err != nil {
		return locationAreaPokemonResp, fmt.Errorf("Error in Unmarshal: %v", err)
	}

	c.cache.Add(url, data)

	return locationAreaPokemonResp, nil
}
