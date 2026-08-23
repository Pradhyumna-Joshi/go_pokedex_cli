package main

import (
	"time"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/cmd"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func main() {
	client := api.NewClient(5*time.Second, 2*time.Minute)
	conf := &config.Config{
		Commands: map[string]config.CliCommand{
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex",
				Callback:    cmd.CommandExit,
			},
			"help": {
				Name:        "help",
				Description: "Displays a help message",
				Callback:    cmd.CommandHelp,
			},
			"map": {
				Name:        "map",
				Description: "Displays the names of 20 location areas in the Pokemon world",
				Callback:    cmd.CommandMapNext,
			},
			"mapb": {
				Name:        "mapb",
				Description: "Displays the previous names of 20 location areas in the Pokemon world",
				Callback:    cmd.CommandMapPrev,
			},
			"explore": {
				Name:        "explore <area_name>",
				Description: "Lists all Pokémons in <area_name>",
				Callback:    cmd.CommandExplore,
			},
			"catch": {
				Name:        "catch <pokemon_name>",
				Description: "Attempt to catch <pokemon_name> and add to the user's Pokedex.",
				Callback:    cmd.CommandCatch,
			},
			"inspect": {
				Name:        "inspect <pokemon_name>",
				Description: "Displays details about <pokemon_name>.",
				Callback:    cmd.CommandInspect,
			},
			"pokedex": {
				Name:        "pokedex",
				Description: "Displays all caught pokemons.",
				Callback:    cmd.CommandPokedex,
			},
		},
		ApiClient: client,
		Pokedex:   map[string]api.Pokemon{},
	}

	startRepl(conf)
}
