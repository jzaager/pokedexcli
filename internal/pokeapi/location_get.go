package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetLocation(areaName string) (Location, error) {
	start := time.Now()
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(areaName); ok {
		pokemonResp := Location{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Location{}, err
		}

		elapsed := time.Since(start)
		logFetchTime(elapsed, true)

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	pokemonResp := Location{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Location{}, err
	}

	elapsed := time.Since(start)
	logFetchTime(elapsed, false)

	c.cache.Add(areaName, data)
	return pokemonResp, nil
}
