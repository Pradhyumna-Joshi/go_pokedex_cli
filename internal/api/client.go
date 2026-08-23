package api

import (
	"net/http"
	"time"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	baseURL    string
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: "https://pokeapi.co/api/v2",
		cache:   pokecache.NewCache(interval),
	}
}
