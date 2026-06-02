package main

import (
	"time"
	"github.com/gclinoz/Pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient:	pokeClient,
	}

	startRepl(cfg)
}
