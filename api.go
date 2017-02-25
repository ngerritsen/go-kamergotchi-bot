package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.kamergotchi.nl/game"

// GameAPI does the network communcation to the kamergotchi API
type GameAPI struct {
	playerToken string
}

// GetGameInfo gets the current state of the game
func (api *GameAPI) GetGameInfo() (Game, error) {
	res, err := apiRequest("/", "GET", api.playerToken, nil)
	if err != nil {
		return Game{}, err
	}

	return parseGame(res)
}

// SpendCare spends care on the kamergotchi
func (api *GameAPI) SpendCare(careType string) (Game, error) {
	body, _ := json.Marshal(map[string]string{"bar": careType})
	res, err := apiRequest("/care", "POST", api.playerToken, body)

	if err != nil {
		return Game{}, err
	}

	return parseGame(res)
}

// ClaimReward claims a reward
func (api *GameAPI) ClaimReward() (Game, error) {
	res, err := apiRequest("/claim", "POST", api.playerToken, nil)

	if err != nil {
		return Game{}, err
	}

	return parseGame(res)
}

func parseGame(res []byte) (Game, error) {
	var info map[string]Game
	if err := json.Unmarshal(res, &info); err != nil {
		return Game{}, err
	}

	return info["game"], nil
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
