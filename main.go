package main

import (
	"time"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/cmd"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/repl"
)

func main() {
	client := api.NewClient(5*time.Second, 2*time.Minute)
	conf := &config.Config{
		Commands: map[string]config.CliCommand{
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex.",
				Callback:    cmd.CommandExit,
			},
			"help": {
				Name:        "help",
				Description: "Display available commands and their descriptions.",
				Callback:    cmd.CommandHelp,
			},
			"map": {
				Name:        "map",
				Description: "Display the next 20 location areas in the Pokémon world.",
				Callback:    cmd.CommandMapNext,
			},
			"mapb": {
				Name:        "mapb",
				Description: "Display the previous 20 location areas in the Pokémon world.",
				Callback:    cmd.CommandMapPrev,
			},
			"explore": {
				Name:        "explore <area_name>",
				Description: "List the Pokémon found in a location area.",
				Callback:    cmd.CommandExplore,
			},
			"catch": {
				Name:        "catch <pokemon_name>",
				Description: "Attempt to catch a Pokémon and add it to your Pokedex.",
				Callback:    cmd.CommandCatch,
			},
			"inspect": {
				Name:        "inspect <pokemon_name>",
				Description: "Display details about a caught Pokémon.",
				Callback:    cmd.CommandInspect,
			},
			"pokedex": {
				Name:        "pokedex",
				Description: "Display all Pokémon in your Pokedex.",
				Callback:    cmd.CommandPokedex,
			},
		},
		ApiClient: client,
		Pokedex:   map[string]api.Pokemon{},
	}

	repl.StartRepl(conf)
}
