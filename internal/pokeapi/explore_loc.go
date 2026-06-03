package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocations(location string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespDeepLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespDeepLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return RespDeepLocations{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}

	locationResp := RespDeepLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespDeepLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
