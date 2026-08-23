package cmd

import (
	"fmt"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandPokedex(c *config.Config, args ...string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("no pokemon's caught!!!")
		return nil
	}
	fmt.Println("Your Pokedex: ")
	for _, pokemon := range c.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil

}
