package main

import (
	"time"
	"github.com/gclinoz/Pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient:	pokeClient,
		pokedex:		make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
