package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	start := time.Now()

	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		elapsed := time.Since(start)
		logFetchTime(elapsed, true)
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	elapsed := time.Since(start)
	logFetchTime(elapsed, false)

	c.cache.Add(url, data)
	return locationsResp, nil
}

func logFetchTime(elapsed time.Duration, cached bool) {
	fmt.Println("======================================")

	if cached {
		fmt.Println("Cache hit!")
	} else {
		fmt.Println("Not cached - fetching from API...")
	}

	fmt.Printf("Request took: %v\n", elapsed)
	fmt.Println("======================================")
}
