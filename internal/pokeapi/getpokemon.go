package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(target string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + target

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokeResp := RespPokeInfo{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return  Pokemon{}, err
	}

	pokem := Pokemon{
		Order: pokeResp.Order,
		Name: pokeResp.Name,
		BaseExperience: pokeResp.BaseExperience,
		Height: pokeResp.Height,
		Weight: pokeResp.Weight,
	}
	return pokem, nil
}
