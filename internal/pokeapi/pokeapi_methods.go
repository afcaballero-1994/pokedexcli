package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)

const baseurl string = "https://pokeapi.co/api/v2"


func GetResources(pageURL *string) (mapResponse, error) {
	url := baseurl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	res, err := http.Get(url)
	var areas mapResponse
	if err != nil {
		fmt.Println(err)
		return areas, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println(res.StatusCode, body)
	}

	if err != nil {
		return areas, err
	}
	
	
	if err = json.Unmarshal(body, &areas); err != nil {
		return areas, err
	}
	
	return areas, nil
}