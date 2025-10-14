package pokeapi

import (
    "net/http"
    "encoding/json"
    "fmt"
    "io"
    "time"
)

const baseurl string = "https://pokeapi.co/api/v2"

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
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        fmt.Println(err)
        return ShallowMapResponse{}, err
    }

    res, err := c.client.Do(req)
    if err != nil {
        return ShallowMapResponse{}, err
    }

    data, err := io.ReadAll(res.Body)
    defer res.Body.Close()

    if res.StatusCode > 299 {
        fmt.Printf("received response with status: %d", res.StatusCode)
    }

    if err != nil {
        return ShallowMapResponse{}, err
    }

    var areas ShallowMapResponse = ShallowMapResponse{}
    if err = json.Unmarshal(data, &areas); err != nil {
        return areas, err
    }

    return areas, nil
}
