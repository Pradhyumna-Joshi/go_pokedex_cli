package config

import (
	"github.com/Pradhyumna-Joshi/go_pokedex/internal/api"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	Commands  map[string]CliCommand
	ApiClient *api.Client
	Pokedex   map[string]api.Pokemon
	NextURL   *string
	PrevURL   *string
}
