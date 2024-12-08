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
	next     string
	previous string
}

func (m *Map) NextLocations() ([]LocationArea, error) {
	if m.next == "" {
		m.next = locationUrl
	}
	resData, err := m.callApi(m.next)
	if err != nil {
		return nil, err
	}
	return resData.Results, nil
}

func (m *Map) PreviousLocations() ([]LocationArea, error) {
	if m.previous == "" {
		return nil, fmt.Errorf("Navigation error: On first page!")
	}
	resData, err := m.callApi(m.previous)
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
	m.next = resData.Next
	if resData.Previous == nil {
		m.previous = ""
	} else {
		m.previous = *(resData.Previous)
	}
	return resData, nil
}
