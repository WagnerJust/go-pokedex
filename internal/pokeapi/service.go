package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"fmt"
	"github.com/WagnerJust/go-pokedex/internal/pokecache"
)
const CACHE_REAP_INTERVAL = 10 * time.Minute

type PokeApiClient struct {
	BaseUrl string
	Client http.Client
	Cache *pokecache.Cache
}

func NewPokeApiClient() *PokeApiClient {
	return &PokeApiClient{
		BaseUrl: "https://pokeapi.co/api/v2",
		Client:  http.Client{},
		Cache: pokecache.NewCache(CACHE_REAP_INTERVAL),
	}
}

func (c *PokeApiClient) makeRequest(req *http.Request, v any) error {
	url := req.URL.String()
	if data, found := c.Cache.Get(url); found {
		fmt.Println("Grabbing data from PokeCache!")
		return json.Unmarshal(data, v)
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	c.Cache.Add(url, data)
	return json.Unmarshal(data, v)

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

	err = c.makeRequest(req, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResponse, nil
}

func (c *PokeApiClient) GetDetailedLocationArea(name string) (DetailedLocationAreaReponse, error) {
	fullUrl := c.BaseUrl + "/location-area/" + name

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return DetailedLocationAreaReponse{}, err
	}

	var detailedLocationArea = DetailedLocationAreaReponse{}

	err = c.makeRequest(req, &detailedLocationArea)
	if err != nil {
		return DetailedLocationAreaReponse{}, err
	}
	return detailedLocationArea, nil
}
