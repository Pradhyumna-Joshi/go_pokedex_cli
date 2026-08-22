package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleApiRequest[T any](c *Client, pageURL *string) (T, error) {

	var zero T

	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return zero, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("HTTP Error : %s", res.Status)
	}

	var resp T
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return zero, err
	}

	return resp, nil

}
