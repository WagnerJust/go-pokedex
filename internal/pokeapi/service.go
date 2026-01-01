package pokeapi

import (
	"encoding/json"
	"net/http"
)

type PokeApiClient struct {
	BaseUrl string
	Client http.Client
}

func NewPokeApiClient() *PokeApiClient {
	return &PokeApiClient{
		BaseUrl: "https://pokeapi.co/api/v2",
		Client:  http.Client{},
	}
}

func (c *PokeApiClient) GetLocationAreas(url *string) (LocationAreaResponse, error) {
	fullUrl := c.BaseUrl + "/location-area"
	if url != nil {
		fullUrl = *url
	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationAreaResponse = LocationAreaResponse{}

	res, err := c.Client.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreaResponse)

	if err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResponse, nil
}
