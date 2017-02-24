package main

import (
	"log"
	"strconv"
	"time"
)

// ClaimLoop recusively keeps claiming rewards
func ClaimLoop(game Game, playerToken string) {
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
	ClaimLoop(game, playerToken)
}
