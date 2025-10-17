package pokeapi

import (
    "net/http"
    "encoding/json"
    "fmt"
    "io"
    "time"
    "github.com/afcaballero-1994/pokedexcli/internal/pokecache"
)

const baseurl string = "https://pokeapi.co/api/v2"
var cache = pokecache.NewCache(5 * time.Minute)

type Client struct {
    client http.Client
}

func NewClient(timeout time.Duration) Client{
    return Client {client: http.Client{
        Timeout: timeout,
        },
    }
}

func (c *Client) GetLocations(pageURL *string) (ShallowMapResponse, error) {
    url := baseurl + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }
    data, exist := cache.Get(url)
    if !exist {
        req, err := http.NewRequest("GET", url, nil)

        if err != nil {
            fmt.Println(err)
            return ShallowMapResponse{}, err
        }

        res, err := c.client.Do(req)
        if err != nil {
            return ShallowMapResponse{}, err
        }

        data, err = io.ReadAll(res.Body)
        defer res.Body.Close()

        if err != nil {
            return ShallowMapResponse{}, err
        }
        cache.Add(url, data)
    }


    var areas ShallowMapResponse = ShallowMapResponse{}
    if err := json.Unmarshal(data, &areas); err != nil {
        return areas, err
    }

    return areas, nil
}

func (c *Client) GetPokemonList(location string) (detailedResponse, error) {
    url := baseurl + "/location-area/" + location

    data, exist := cache.Get(url)
    if !exist {
        req, err := http.NewRequest("GET", url, nil)

        if err != nil {
            fmt.Println(err)
            return detailedResponse{}, err
        }

        res, err := c.client.Do(req)
        if err != nil {
            return detailedResponse{}, err
        }

        data, err = io.ReadAll(res.Body)
        defer res.Body.Close()

        if err != nil {
            return detailedResponse{}, err
        }
        cache.Add(url, data)
    }


    var detail detailedResponse = detailedResponse{}
    if err := json.Unmarshal(data, &detail); err != nil {
        return detail, err
    }

    return detail, nil
}

func (c *Client) GetPokemon(location string) (Pokemon, error) {
    url := baseurl + "/pokemon/" + location

    data, exist := cache.Get(url)
    if !exist {
        req, err := http.NewRequest("GET", url, nil)

        if err != nil {
            fmt.Println(err)
            return Pokemon{}, err
        }

        res, err := c.client.Do(req)
        if err != nil {
            return Pokemon{}, err
        }

        data, err = io.ReadAll(res.Body)
        defer res.Body.Close()

        if err != nil {
            return Pokemon{}, err
        }
        cache.Add(url, data)
    }


    var poke Pokemon = Pokemon{}
    if err := json.Unmarshal(data, &poke); err != nil {
        return poke, err
    }

    return poke, nil
}
