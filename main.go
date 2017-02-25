package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Go Kamergotchi bot started.")

	playerToken := getPlayerToken()
	api := &GameAPI{playerToken}
	errChan := make(chan error)
	game, err := api.GetGameInfo()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Retrieved player info for kamergotchi " + game.Gotchi.getInfo() + ".")

	go ClaimLoop(game, api, errChan)
	go CareLoop(game, api, errChan)

	log.Fatal(<-errChan)
}

func getPlayerToken() string {
	if len(os.Args) < 2 {
		log.Fatal("Player token not provided, please provide the token as command line argument.")
	}

	return os.Args[1]
}
