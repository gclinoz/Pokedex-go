package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		if c, exists := commands[input]; exists {
			c.callback()
		} else {
			fmt.Println("Unknown command:", input)
		}
		fmt.Print("Pokedex > ")
	}
}

type cliCommand struct {
	name	string
	description	string
	callback	func() error
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpMsg() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}
