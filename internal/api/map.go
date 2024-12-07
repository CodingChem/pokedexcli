package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationUrl string = "https://pokeapi.co/api/v2/location/"

type ApiResponse struct {
	Previous *string        `json:"previous"`
	Next     string         `json:"next"`
	Results  []LocationArea `json:"results"`
}
type LocationArea struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Map struct {
	Next     string
	Previous string
}

func (m *Map) NextLocations() ([]LocationArea, error) {
	if m.Next == "" {
		m.Next = locationUrl
	}
	resData, err := m.callApi(m.Next)
	if err != nil {
		return nil, err
	}
	return resData.Results, nil
}

func (m *Map) PreviousLocations() ([]LocationArea, error) {
	if m.Previous == "" {
		return nil, fmt.Errorf("Navigation error: On first page!")
	}
	resData, err := m.callApi(m.Previous)
	if err != nil {
		return nil, err
	}
	return resData.Results, nil
}

func (m *Map) callApi(url string) (ApiResponse, error) {
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
	m.Next = resData.Next
	if resData.Previous == nil {
		m.Previous = ""
	} else {
		m.Previous = *(resData.Previous)
	}
	return resData, nil
}
