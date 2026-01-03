package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
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
func (c *PokeApiClient) GetLocationAreas(url *string) (LocationAreaList, error) {
	fullUrl := c.BaseUrl + "/location-area"
	if url != nil {
		fullUrl = *url
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	var locationAreaResponse = LocationAreaList{}

	err = c.makeRequest(req, &locationAreaResponse)
	if err != nil {
		return LocationAreaList{}, err
	}
	return locationAreaResponse, nil
}

func (c *PokeApiClient) GetDetailedLocationArea(name string) (LocationAreaSingle, error) {
	fullUrl := c.BaseUrl + "/location-area/" + name

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaSingle{}, err
	}

	var detailedLocationArea = LocationAreaSingle{}

	err = c.makeRequest(req, &detailedLocationArea)
	if err != nil {
		return LocationAreaSingle{}, err
	}
	return detailedLocationArea, nil
}

func (c *PokeApiClient) GetPokemon(name string) (Pokemon, error) {
	fullUrl := c.BaseUrl + "/pokemon/" + name

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon = Pokemon{}

	err = c.makeRequest(req, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
