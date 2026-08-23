package cmd

import (
	"errors"
	"fmt"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandExplore(c *config.Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a location name")
	}
	location := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	resp, err := api.HandleApiRequest[api.AreaInfo](c.ApiClient, &url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
