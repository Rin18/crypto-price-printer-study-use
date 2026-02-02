package main

import (
	"fmt"
	"io"
	"net/http"
)

func get_response(endpoint string) (string, error) {
	baseURL := "https://fapi.binance.com"
	fullURL := baseURL + endpoint
	resp, err := http.Get(fullURL)
	// check error
	if err != nil {
		return "", err
	}
	//close connection
	defer resp.Body.Close()
	//check error
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Status error:", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	//check error
	if err != nil {
		return "", err
	}

	return string(body), err
}

func main() {
	response, err := get_response("/fapi/v1/ping")
	//check error
	if err != nil {
		fmt.Println("Request failed, error:", err)
	}
	fmt.Println(response)
}
