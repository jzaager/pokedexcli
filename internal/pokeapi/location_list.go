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

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedLocations, ok := c.getCached(url)
	if ok {
		fmt.Println()
		fmt.Println("Cache hit!")
		elapsed := time.Since(start)
		fmt.Printf("Request took: %v\n", elapsed)
		fmt.Println()
		return cachedLocations, nil
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
	c.cache.Add(url, data)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println("Not cached - fetching from API...")
	fmt.Printf("Request took: %v\n", elapsed)
	fmt.Println()
	return locationsResp, nil
}

func (c *Client) getCached(url string) (RespShallowLocations, bool) {
	cacheData, ok := c.cache.Get(url)
	if !ok {
		return RespShallowLocations{}, false
	}

	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(cacheData, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, false
	}

	return locationsResp, true
}
