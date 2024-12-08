package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationUrl string = "https://pokeapi.co/api/v2/location/"

type ApiResponse struct {
	Previous *string         `json:"previous"`
	Next     string          `json:"next"`
	Results  json.RawMessage `json:"results"`
}

func NextLocations(url string) (ApiResponse, error) {
	if url == "" {
		url = locationUrl
	}
	resData, err := callApi(url)
	if err != nil {
		return ApiResponse{}, err
	}
	return resData, nil
}

func callApi(url string) (ApiResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return ApiResponse{}, fmt.Errorf("Invalid error code: %d", res.StatusCode)
	}

	dec := json.NewDecoder(res.Body)
	var resData ApiResponse
	dec.Decode(&resData)

	return resData, nil
}
