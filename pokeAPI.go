package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeCities struct {
	next     string
	previous string
	results  []struct {
		name string
		url  string
	}
}

func requestPokeAPI(url string) (pokeCities, error) {
	res, err := http.Get(url)
	if err != nil {
		return pokeCities{}, fmt.Errorf("Couldn't fetch data from %v.\n%w", url, err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeCities{}, fmt.Errorf("Couldn't read from body.\n%w", err)
	}
	poke := pokeCities{}
	if err = json.Unmarshal(data, &poke); err != nil {
		return pokeCities{}, err
	}
	defer res.Body.Close()
	return poke, nil

}

func mapn() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	ret, err := requestPokeAPI(url)
	if err != nil {
		return err
	}
	fmt.Println(ret)

	return nil
}
