package cmd

import (
	"fmt"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandHelp(c *config.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for key, cmd := range c.Commands {
		fmt.Println(key + ": " + cmd.Description)
	}
	return nil
}
