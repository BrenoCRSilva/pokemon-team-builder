package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/BrenoCRSilva/pokemon-team-builder/cache"
)

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		baseUrl: "https://pokeapi.co/api/v2",
		cache:   cache.NewCache(cacheInterval),
		client:  &http.Client{},
	}
}

func (c *Client) fetchFromCacheOrAPI(url string) ([]byte, error) {
	if cached, ok := c.cache.Get(url); ok {
		return cached, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	c.cache.Add(url, data)
	return data, nil
}

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	var pokemon Pokemon
	url := fmt.Sprintf("%s/pokemon/%s", c.baseUrl, name)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	for i, pokemonType := range pokemon.Types {
		typeURL := fmt.Sprintf("%s/type/%s", c.baseUrl, pokemonType.Details.Name)
		typeDetails, err := c.fetchTypeDetails(typeURL)
		if err != nil {
			log.Printf("Error fetching type details for %s: %v", pokemonType.Details.Name, err)
			continue
		}
		pokemon.Types[i].Details = *typeDetails
	}

	return pokemon, nil
}

func (c *Client) fetchTypeDetails(typeURL string) (*TypeDetail, error) {
	data, err := c.fetchFromCacheOrAPI(typeURL)
	if err != nil {
		return nil, err
	}

	var typeDetails TypeDetail
	err = json.Unmarshal(data, &typeDetails)
	if err != nil {
		return nil, err
	}

	return &typeDetails, nil
}
