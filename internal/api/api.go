package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleApiRequest[T any](c *Client, pageURL *string) (T, error) {

	var zero T

	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte

	entry, exists := c.cache.Get(url)
	if exists {
		fmt.Println("FROM CACHE")
		fmt.Println()
		data = entry
	} else {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return zero, err
		}

		response, err := c.httpClient.Do(req)
		if err != nil {
			return zero, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return zero, fmt.Errorf("HTTP Error : %s", response.Status)
		}

		data, err = io.ReadAll(response.Body)
		if err != nil {
			return zero, err
		}
		c.cache.Add(url, data)
	}

	var resp T
	if err := json.Unmarshal(data, &resp); err != nil {
		return zero, err
	}

	return resp, nil

}
