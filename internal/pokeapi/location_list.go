package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		if err := json.Unmarshal(val, &locationsResp); err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	if err = json.Unmarshal(dat, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
