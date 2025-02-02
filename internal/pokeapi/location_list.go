package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RespLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.fetchAPIData(url)
	if err != nil {
		return RespLocations{}, err
	}
	fmt.Printf("data: %s\n", data)

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) fetchAPIData(url string) ([]byte, error) {
	if data, found := c.cache.Get(url); found {
		log.Println("Data from cache!")
		return data, nil
	} else {
		log.Println("Data missed")
	}

	req, err := http.NewRequest("GET", url, nil)
	log.Println("Data from web!")
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)
	log.Println("Data added")

	return data, nil

}
