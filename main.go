package main

import (
    "time"
    "github.com/afcaballero-1994/pokedexcli/internal/pokeapi"
)

func main() {
    pclient := pokeapi.NewClient(2 * time.Second)
    c := config {
        pokeapiClient: pclient,
        caughtPokemon: map[string]pokeapi.Pokemon{},
    }
    startRepl(&c)
}
