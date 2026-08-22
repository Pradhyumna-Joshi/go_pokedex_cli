package api

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	baseURL    string
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}
