package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationAreaResp LocationAreaResp
	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &locationAreaResp); err != nil {
			return locationAreaResp, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResp, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResp, err
	}

	if err := json.Unmarshal(data, &locationAreaResp); err != nil {
		return locationAreaResp, err
	}

	c.cache.Add(url, data)
	return locationAreaResp, nil
}
