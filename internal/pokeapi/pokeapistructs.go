package pokeapi

type mapResponse struct {
	Count    int      `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
