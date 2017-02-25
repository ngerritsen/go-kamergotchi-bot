package main

import (
	"log"
	"strconv"
	"time"
)

// CareLoop recusively keeps caring for the gotchi
func CareLoop(game Game, api *GameAPI, errChan chan<- error) {
	if game.CareLeft == 0 {
		reset, err := time.Parse(time.RFC3339, game.CareReset)
		if err != nil {
			errChan <- err
			return
		}

		duration := calcCareWaitDuration(reset)
		log.Println("Cannot spend any more care, waiting for " + duration.String() + "...")
		time.Sleep(duration)

		game, err = api.GetGameInfo()
		if err != nil {
			errChan <- err
			return
		}
	}

	careTypeToGive := determineCareTypeToGive(game.Stats)
	game, err := api.SpendCare(careTypeToGive)
	if err != nil {
		errChan <- err
		return
	}

	log.Println("Spent care on " + careTypeToGive + " new score: " + strconv.Itoa(game.Score) + ".")
	CareLoop(game, api, errChan)
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

func calcCareWaitDuration(reset time.Time) time.Duration {
	timeToWait := reset.Sub(time.Now())
	return timeToWait + randomSeconds(60)
}
