package config

import "github.com/Pradhyumna-Joshi/go_pokedex/internal/api"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Commands   map[string]CliCommand
	HttpClient *api.Client
	NextURL    *string
	PrevURL    *string
}
