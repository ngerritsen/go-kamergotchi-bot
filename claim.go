package main

import (
	"log"
	"strconv"
	"time"
)

// ClaimLoop recusively keeps claiming rewards
func ClaimLoop(game Game, api *GameAPI, errChan chan<- error) {
	resetTime, err := time.Parse(time.RFC3339, game.ClaimReset)
	if err != nil {
		errChan <- err
		return
	}

	duration := calcClaimWaitDuration(resetTime)

	if duration > 0 {
		log.Println("Making next claim in " + duration.String() + "...")
		time.Sleep(duration)
	}

	game, err = api.ClaimReward()
	if err != nil {
		errChan <- err
		return
	}

	log.Println("Claimed reward, new score: " + strconv.Itoa(game.Score) + ".")
	ClaimLoop(game, api, errChan)
}

func calcClaimWaitDuration(resetTime time.Time) time.Duration {
	timeToWait := resetTime.Sub(time.Now())
	return timeToWait + 80*time.Second + randomSeconds(60)
}
