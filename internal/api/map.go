package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	locationUrl     string = "https://pokeapi.co/api/v2/location/"
	locationAreaUrl string = "https://pokeapi.co/api/v2/location-area/"
	pokemonUrl      string = "https://pokeapi.co/api/v2/pokemon/"
)

type ApiResponse struct {
	Previous *string         `json:"previous"`
	Next     string          `json:"next"`
	Results  json.RawMessage `json:"results"`
}

func GetLocations(url string) ([]byte, error) {
	if url == "" {
		url = locationAreaUrl
	}
	res, err := callApi(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetPokemon(pokemon string) ([]byte, error) {
	res, err := callApi(pokemonUrl + pokemon)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetLocation(name string) ([]byte, error) {
	url := locationAreaUrl + name
	res, err := callApi(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func callApi(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid error code: %d", res.StatusCode)
	}

	return io.ReadAll(res.Body)
}
