package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	log.Println("Go Kamergotchi bot started.")
	playerToken := getPlayerToken()
	game := GetGameInfo(playerToken)
	log.Println("Retrieved player info for kamergotchi " + game.Gotchi.getInfo() + ".")

	go claimLoop(game, playerToken)
	careLoop(game, playerToken)
}

func claimLoop(game Game, playerToken string) {
	reset, err := time.Parse(time.RFC3339, game.ClaimReset)
	if err != nil {
		log.Fatal(err)
	}

	duration := GetClaimWaitDuration(reset)

	if duration > time.Second {
		log.Println("Making next claim in " + duration.String() + "...")
		time.Sleep(duration)
	}

	game = ClaimReward(playerToken)
	log.Println("Claimed reward, new score: " + strconv.Itoa(game.Score) + ".")
	claimLoop(game, playerToken)
}

func careLoop(game Game, playerToken string) {
	if game.CareLeft == 0 {
		reset, err := time.Parse(time.RFC3339, game.CareReset)
		if err != nil {
			log.Fatal(err)
		}

		duration := GetCareWaitDuration(reset)
		log.Println("Cannot spend any more care, waiting for " + duration.String() + "...")
		time.Sleep(duration)
		game = GetGameInfo(playerToken)
	}

	careTypeToGive := determineCareTypeToGive(game.Stats)
	game = SpendCare(careTypeToGive, playerToken)
	log.Println("Spent care on " + careTypeToGive + " new score: " + strconv.Itoa(game.Score) + ".")
	careLoop(game, playerToken)
}

func determineCareTypeToGive(stats map[string]int) string {
	var careTypeToGive string
	max := 101

	for careType, value := range stats {
		if value < max {
			max = value
			careTypeToGive = careType
		}
	}

	return careTypeToGive
}

func getPlayerToken() string {
	if len(os.Args) < 2 {
		log.Fatal("Player token not provided, please provide the token as command line argument.")
	}

	return os.Args[1]
}
