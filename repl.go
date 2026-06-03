package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/gclinoz/Pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient	pokeapi.Client
	pokedex			map[string]pokeapi.Pokemon
	nextLocURL		*string
	prevLocURL		*string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		haveArg := len(words) > 1

		command, exists := getCommands()[commandName]
		if exists && haveArg {
			err := command.callback(cfg, words[1])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	clean_string := []string{}

	trimmed := strings.TrimSpace(text)

	for _, s := range strings.Split(trimmed, " ") {
		if s != "" {
			clean_string = append(clean_string, strings.ToLower(s))
		}
	}

	return clean_string
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
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
		"explore": {
			name:	"explore",
			description:	"Show all the Pokemon located there",
			callback:	explorePoke,
		},
		"catch": {
			name:	"catch",
			description:	"Catch your favorite Pokemon",
			callback:	catchPoke,
		},
		"inspect": {
			name:	"inspect",
			description:	"Displays Pokemon information",
			callback:	showPoke,
		},
	}
}
