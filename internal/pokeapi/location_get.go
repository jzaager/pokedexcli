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

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		elapsed := time.Since(start)
		logFetchTime(elapsed, true)

		return locationResp, nil
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

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	elapsed := time.Since(start)
	logFetchTime(elapsed, false)

	c.cache.Add(url, data)
	return locationResp, nil
}
