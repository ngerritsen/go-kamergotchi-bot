package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Go Kamergotchi bot started.")
	playerToken := getPlayerToken()
	game := GetGameInfo(playerToken)
	log.Println("Retrieved player info for kamergotchi " + game.Gotchi.getInfo() + ".")

	go ClaimLoop(game, playerToken)
	go CareLoop(game, playerToken)

	select {} // Prevent program from exiting
}

func getPlayerToken() string {
	if len(os.Args) < 2 {
		log.Fatal("Player token not provided, please provide the token as command line argument.")
	}

	return os.Args[1]
}
