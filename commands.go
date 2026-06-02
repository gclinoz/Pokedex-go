package main

import (
	"fmt"
	"os"
	"errors"
)

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
	res, err := cfg.pokeapiClient.ExploreLocations(site[0])
	if err != nil {
		return err
	}

	for _, loc := range res.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil
}
