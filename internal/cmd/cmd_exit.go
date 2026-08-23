package cmd

import (
	"fmt"
	"os"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func CommandExit(c *config.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
