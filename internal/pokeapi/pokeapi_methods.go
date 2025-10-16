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

func (c *Client) GetResources(pageURL *string) (ShallowMapResponse, error) {
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

        if res.StatusCode > 299 {
            fmt.Printf("received response with status: %d", res.StatusCode)
        }

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
