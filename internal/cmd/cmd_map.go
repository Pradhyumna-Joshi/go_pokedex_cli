package cmd

import (
	"errors"
	"fmt"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandMapNext(c *config.Config, args ...string) error {

	resp, err := api.HandleApiRequest[api.Locations](c.ApiClient, c.NextURL)
	if err != nil {
		return err
	}

	c.NextURL = resp.Next
	c.PrevURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func CommandMapPrev(c *config.Config, args ...string) error {

	if c.PrevURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := api.HandleApiRequest[api.Locations](c.ApiClient, c.PrevURL)
	if err != nil {
		return err
	}

	c.NextURL = resp.Next
	c.PrevURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
