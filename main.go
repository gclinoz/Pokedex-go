package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"encoding/json"
	"net/http"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := config{
		Next:	"https://pokeapi.co/api/v2/location-area/",
		Previous:	nil,
	}

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		if c, exists := commands[input]; exists {
			c.callback(&cfg)
		} else {
			fmt.Println("Unknown command:", input)
		}
		fmt.Print("Pokedex > ")
	}
}

type cliCommand struct {
	name	string
	description	string
	callback	func(*config) error
}

type config struct {
	Next	string
	Previous	*string
}

type Location struct {
	Count	int	`json:"count"`
	Next	string	`json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

func getCommands() map[string]cliCommand {		
	return map[string]cliCommand{
		"exit": {
			name:	"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name:	"help",
			description:	"Displays a help message",
			callback:	helpMsg,
		},
		"map": {
			name:	"map",
			description:	"Show the names of 20 location areas",
			callback:	getMap,
		},
		"mapb": {
			name:	"mapb",
			description:	"Show previous page of the location areas",
			callback:	getMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpMsg(cfg *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}

func getMap(cfg *config) error {
	res, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	location := Location{}
	if err := json.Unmarshal(data, &location); err != nil {
		return err
	}
	cfg.Next = location.Next
	cfg.Previous = location.Previous

	for _, res := range location.Results {
		fmt.Println(res.Name)
	}

	return nil
}

func getMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(*cfg.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	location := Location{}
	if err := json.Unmarshal(data, &location); err != nil {
		return err
	}
	cfg.Next = location.Next
	cfg.Previous = location.Previous

	for _, res := range location.Results {
		fmt.Println(res.Name)
	}

	return nil
}
