package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	start := time.Now()
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		elapsed := time.Since(start)
		logFetchTime(elapsed, true)

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	elapsed := time.Since(start)
	logFetchTime(elapsed, false)

	c.cache.Add(url, data)
	return pokemonResp, nil
}
