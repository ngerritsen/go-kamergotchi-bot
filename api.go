package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const apiURL = "https://api.kamergotchi.nl/game"

// GetGameInfo gets the current state of the game
func GetGameInfo(playerToken string) Game {
	res, err := apiRequest(playerToken)
	if err != nil {
		log.Fatal(err)
	}

	var info map[string]Game
	if err := json.Unmarshal(res, &info); err != nil {
		log.Fatal(err)
	}

	return info["game"]
}

func apiRequest(playerToken string) ([]byte, error) {
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
