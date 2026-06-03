package main

import (
	"fmt"
	"os"
	"errors"
	"math/rand"
)

const catchThres = 10

func commandExit(cfg *config, site ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpMsg(cfg *config, site ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func getMap(cfg *config, site ...string) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = res.Next
	cfg.prevLocURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func getMapb(cfg *config, site ...string) error {
	if cfg.prevLocURL == nil {
		return errors.New("you're on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocations(cfg.prevLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = res.Next
	cfg.prevLocURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func explorePoke(cfg *config, site ...string) error {
	if len(site) == 0 {
		fmt.Println("Please provide location to explore")
		return nil
	}

	res, err := cfg.pokeapiClient.ExploreLocations(site[0])
	if err != nil {
		return err
	}

	for _, loc := range res.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil
}

func catchPoke(cfg *config, target ...string) error {
	if len(target) == 0 {
		fmt.Println("Please provide target to catch")
		return nil
	}

	res, err := cfg.pokeapiClient.GetPokemon(target[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", target[0])
	result := rand.Intn(res.BaseExperience)
	if result > catchThres {
		fmt.Println(target[0], "escaped!")
	} else {
		fmt.Println(target[0], "was caught!")
		cfg.pokedex[target[0]] = res
	}
	return nil
}
