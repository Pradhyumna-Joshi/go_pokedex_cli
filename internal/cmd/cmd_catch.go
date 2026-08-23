package cmd

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandCatch(c *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}
	pokemon := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	resp, err := api.HandleApiRequest[api.Pokemon](c.ApiClient, &url)
	if err != nil {
		return err
	}
	exp := resp.BaseExperience
	catch := rand.IntN(exp)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if catch > int(0.5*float32(exp)) {
		c.Pokedex[pokemon] = resp
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
