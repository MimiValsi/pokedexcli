package pokeapi

import (
	"encoding/json"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	data, err := c.fetchAPIData(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
