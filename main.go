package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Go Kamergotchi bot started.")

	playerToken := getPlayerToken()
	api := &GameAPI{playerToken}
	game, err := api.GetGameInfo()
	if err != nil {
		log.Fatal(err)
	}

	logInfo(game)

	go ClaimLoop(game, api)
	go CareLoop(game, api)

	select {} // Prevent application from exiting
}

func logInfo(game Game) {
	score := strconv.Itoa(game.Score)
	msg := "Retrieved player info for kamergotchi " + game.Gotchi.getInfo() + ", current score: " + score + "."

	log.Println(msg)
}

func getPlayerToken() string {
	if len(os.Args) < 2 {
		log.Fatal("Player token not provided, please provide the token as command line argument.")
	}

	return os.Args[1]
}
