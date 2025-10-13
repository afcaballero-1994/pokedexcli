package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/afcaballero-1994/pokedexcli/internal/pokeapi"
)


func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	message := `
Welcome to the Pokedex!
Usage:


help: Display a help message
exit: Exit the Pokedex
map:  Display the next 20 areas
mapb: Display the Previous 20 areas
	`
	fmt.Println(message)
	return nil
}

func commandMap(c * config) error {
	area, err := pokeapi.GetResources(c.Next)
	if err != nil {
		return nil
	}
	c.Next = area.Next
	c.Previous = area.Previous
	for _, a := range area.Results {
		fmt.Println(a.Name)
	}
	return err
}

func commandMapb(c * config) error {
	if c.Previous == nil {
		fmt.Println("First Page")
		return nil
	}
	area, err := pokeapi.GetResources(c.Previous)
	if err != nil {
		return nil
	}
	c.Next = area.Next
	c.Previous = area.Previous
	for _, a := range area.Results {
		fmt.Println(a.Name)
	}
	return err
}

type command struct {
	name string
	description string
	callback func(c *config) error
}

type config struct {
	Next *string
	Previous *string
}

func cleanInput(text string) []string {
	out := strings.ToLower(text)
	result := strings.Fields(out)
	return result
}

func startRepl() {
	commandRegister := map[string]command {
		"exit": {
			name: "Exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Show a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Show next 20 areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show Previous 20 areas",
			callback: commandMapb,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	var c config
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			break
		}
		command, exist := commandRegister[cleaned[0]]
		if !exist {
			fmt.Println("Unknown command")
			continue
		}
		
		err := command.callback(&c)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}