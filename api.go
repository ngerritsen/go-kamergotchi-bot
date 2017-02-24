package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const apiURL = "https://api.kamergotchi.nl/game"

// GetGameInfo gets the current state of the game
func GetGameInfo(playerToken string) Game {
	res, err := apiRequest("/", "GET", playerToken, nil)
	if err != nil {
		log.Fatal(err)
	}

	return parseGame(res)
}

// SpendCare spends care on the kamergotchi
func SpendCare(careType string, playerToken string) Game {
	body, _ := json.Marshal(map[string]string{"bar": careType})
	res, err := apiRequest("/care", "POST", playerToken, body)

	if err != nil {
		log.Fatal(err)
	}

	return parseGame(res)
}

// ClaimReward claims a reward
func ClaimReward(playerToken string) Game {
	res, err := apiRequest("/claim", "POST", playerToken, nil)

	if err != nil {
		log.Fatal(err)
	}

	return parseGame(res)
}

func parseGame(res []byte) Game {
	var info map[string]Game
	if err := json.Unmarshal(res, &info); err != nil {
		log.Fatal(err)
	}

	return info["game"]
}

func apiRequest(path string, method string, playerToken string, body []byte) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, apiURL+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-player-token", playerToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resBody, err
	}

	return resBody, nil
}
