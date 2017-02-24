package main

import (
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.kamergotchi.nl/game"

// APIRequest does an api request to kamergotchi
func APIRequest(playerToken string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-player-token", playerToken)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}
