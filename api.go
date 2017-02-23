package main

import (
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.kamergotchi.nl/game"

// ApiRequest does an api request to kamergotchi
func ApiRequest(playerToken string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("x-player-token", playerToken)

	if err != nil {
		return "", err
	}

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
