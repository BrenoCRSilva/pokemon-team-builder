package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BrenoCRSilva/pokemon-team-builder/cache"
)

type Client struct {
	baseUrl string
	cache   *cache.Cache
	client  *http.Client
}

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
	return pokemon, nil
}
