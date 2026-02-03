package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Ticker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"markPrice"`
	Time   int64  `json:"time"`
}

var userTicker Ticker

func get_response(endpoint string, params map[string]string) ([]byte, error) {
	baseURL := "https://fapi.binance.com"
	primaryURL := baseURL + endpoint
	var fullURL string
	// process with URL query parameters
	if params == nil {
		fullURL = primaryURL
	} else {
		u, err := url.Parse(primaryURL)
		if err != nil {
			return nil, err
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
		return nil, err
	}
	//close connection
	defer resp.Body.Close()
	//check error
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	//check error
	if err != nil {
		return nil, err
	}

	return body, err
}

func main() {
	response, err := get_response("/fapi/v1/premiumIndex", map[string]string{"symbol": "BTCUSDT"})
	//check error
	if err != nil {
		fmt.Println("Request failed, error:", err)
	} else {
		err := json.Unmarshal(response, &userTicker)
		if err != nil {
			fmt.Println("unmarshal failed, error:", err)
		} else {
			fmt.Println(userTicker)
		}
	}
}
