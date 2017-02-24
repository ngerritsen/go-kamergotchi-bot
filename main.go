package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	log.Println("Go Kamergotchi bot started.")

	playerToken := os.Args[1]

	game := getGameInfo(playerToken)
	log.Println(game.Gotchi.getInfo())
}

func getGameInfo(playerToken string) Game {
	res, err := APIRequest(playerToken)
	if err != nil {
		log.Fatal(err)
	}

	var info map[string]Game
	if err := json.Unmarshal(res, &info); err != nil {
		log.Fatal(err)
	}

	return info["game"]
}
