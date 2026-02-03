package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func get_response(endpoint string, params map[string]string) (string, error) {
	baseURL := "https://fapi.binance.com"
	primaryURL := baseURL + endpoint
	var fullURL string
	// process with URL query parameters
	if params == nil {
		fullURL = primaryURL
	} else {
		u, err := url.Parse(primaryURL)
		if err != nil {
			return "", err
		} else {
			q := u.Query()
			for key, value := range params {
				q.Add(key, value)
			}
			u.RawQuery = q.Encode()
			fullURL = u.String()
		}
	}

	resp, err := http.Get(fullURL)
	// check error
	if err != nil {
		return "", err
	}
	//close connection
	defer resp.Body.Close()
	//check error
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Status error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	//check error
	if err != nil {
		return "", err
	}

	return string(body), err
}

func main() {
	response, err := get_response("/fapi/v1/premiumIndex", nil)
	//check error
	if err != nil {
		fmt.Println("Request failed, error:", err)
	} else {
		fmt.Println(response)
	}
}
