package pokeapi

// response from fetching and decoding from pokeapi
// https://pokeapi.co/api/v2/location-area
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
