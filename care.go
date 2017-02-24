package main

import (
	"log"
	"strconv"
	"time"
)

// CareLoop recusively keeps caring for the gotchi
func CareLoop(game Game, playerToken string) {
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
	CareLoop(game, playerToken)
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
