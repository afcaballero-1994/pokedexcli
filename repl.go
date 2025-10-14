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
    fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
    for _, command := range getCommands() {
        fmt.Printf("%s: %s\n", command.name, command.description)
    }
    fmt.Println()
    return nil
}

func commandMap(c * config) error {
    area, err := c.pokeapiClient.GetResources(c.Next)
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
    area, err := c.pokeapiClient.GetResources(c.Previous)
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
    pokeapiClient pokeapi.Client
    Next *string
    Previous *string
}

func cleanInput(text string) []string {
    out := strings.ToLower(text)
    result := strings.Fields(out)
    return result
}

func getCommands() map[string]command{
    return map[string]command {
        "exit": {
            name: "exit",
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
            description: "Show next set of locations",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb",
            description: "Show Previous set of locations",
            callback: commandMapb,
        },
    }
}

func startRepl(c *config) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        cleaned := cleanInput(scanner.Text())
        if len(cleaned) == 0 {
            break
        }
        command, exist := getCommands()[cleaned[0]]
        if !exist {
            fmt.Println("Unknown command")
            continue
        }

        err := command.callback(c)
        if err != nil {
            fmt.Println(err)
            continue
        }

    }
}
