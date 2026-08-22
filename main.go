package main

import (
	"time"

	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/cmd"
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/config"
)

func main() {
	client := api.NewClient(3 * time.Second)
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
		},
		HttpClient: client,
	}

	startRepl(conf)
}
