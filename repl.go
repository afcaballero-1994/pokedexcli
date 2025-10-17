package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "errors"
    "time"
    "math/rand"
    "github.com/afcaballero-1994/pokedexcli/internal/pokeapi"
)

type command struct {
    name string
    description string
    callback func(c *config, args ...string) error
}

type config struct {
    pokeapiClient pokeapi.Client
    Next          *string
    Previous      *string
    caughtPokemon map[string]pokeapi.Pokemon
}


func commandExit(c *config, args ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(c *config, args ...string) error {
    fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
    for _, command := range getCommands() {
        fmt.Printf("%s: %s\n", command.name, command.description)
    }
    fmt.Println()
    return nil
}

func commandMap(c * config, args ...string) error {
    area, err := c.pokeapiClient.GetLocations(c.Next)
    if err != nil {
        return err
    }
    c.Next = area.Next
    c.Previous = area.Previous
    for _, a := range area.Results {
        fmt.Println(a.Name)
    }
    return nil
}

func commandMapb(c * config, args ...string) error {
    if c.Previous == nil {
        return errors.New("You are on the first page")
    }
    area, err := c.pokeapiClient.GetLocations(c.Previous)
    if err != nil {
        return err
    }
    c.Next = area.Next
    c.Previous = area.Previous
    for _, a := range area.Results {
        fmt.Println(a.Name)
    }
    return nil
}

func commandExplore(c * config, args ...string) error {
    if len(args) != 1 {
        return errors.New("You need to provide exactly one argument")
    }

    area, err := c.pokeapiClient.GetPokemonList(args[0])
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s...", args[0])
    fmt.Println("Found Pokemon:")
    for _, p := range area.PokemonEncounters {
        fmt.Println("  -", p.Pokemon.Name)
    }
    return nil
}

func commandCatch(c *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("You need to provide exactly one argument")
    }
    name := args[0]
    fmt.Printf("Throwing a Pokeball at %s...\n", name)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    poke, err := c.pokeapiClient.GetPokemon(name)
    if err != nil {
        return err
    }

    chance := r.Intn(poke.BaseExperience + 120)
    if chance <= poke.BaseExperience - 5 {
        return fmt.Errorf("%s escaped!", name)
    } else {
        c.caughtPokemon[poke.Name] = poke
        fmt.Printf("%s was caught!\n", name)
        return nil
    }
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
        "explore": {
            name: "explore",
            description: "List all Pokemon located in the area",
            callback: commandExplore,
        },
        "catch": {
            name: "catch",
            description: "Attempt to get a Pokemon",
            callback: commandCatch,
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

        args := []string{}

        if len(cleaned) > 1 {
            args = cleaned[1:]
        }
        err := command.callback(c, args...)
        if err != nil {
            fmt.Println(err)
            continue
        }

    }
}
