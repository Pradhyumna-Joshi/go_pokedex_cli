package cmd

import (
	"errors"
	"fmt"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandInspect(c *config.Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	pokemon := args[0]

	val, exists := c.Pokedex[pokemon]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	displayPokemon(&val)

	return nil
}

func displayPokemon(p *api.Pokemon) {
	fmt.Println("Name: ", p.Name)
	fmt.Println("Height: ", p.Height)

	fmt.Println("Weight: ", p.Weight)
	fmt.Println("Stats: ")
	for _, stat := range p.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range p.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

}
