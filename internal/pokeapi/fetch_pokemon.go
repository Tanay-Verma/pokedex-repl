package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	var pokemonResp Pokemon

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &pokemonResp); err != nil {
			return pokemonResp, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonResp, fmt.Errorf("Error in creating NewRequest: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonResp, fmt.Errorf("Error in calling the url: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return pokemonResp, fmt.Errorf("%v in fetching the pokemon. Please enter a valid pokemon", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonResp, fmt.Errorf("Error in loading the data to memory: %v", err)
	}

	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return pokemonResp, fmt.Errorf("Error in Unmarshal: %v", err)
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}
