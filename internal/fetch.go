package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type JsonData struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

func FetchLocationData(url string) (JsonData, error) {
	res, err := http.Get(url)
	if err != nil {
		return JsonData{}, fmt.Errorf("Error fetching data from API: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return JsonData{}, fmt.Errorf("Fetch request failed with status code: %d\nStatus: %s", res.StatusCode, res.Status)
	}

	var jsonData JsonData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&jsonData)

	if err != nil {
		return JsonData{}, fmt.Errorf("Error decoding json response: %w", err)
	}

	return jsonData, nil
}
