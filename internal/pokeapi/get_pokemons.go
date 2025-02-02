package pokeapi

import "encoding/json"

func (c *Client) LocationPokemons(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	data, err := c.fetchAPIData(url)
	if err != nil {
		return Location{}, err
	}

	zone := Location{}
	err = json.Unmarshal(data, &zone)
	if err != nil {
		return Location{}, err
	}

	return zone, nil
}
