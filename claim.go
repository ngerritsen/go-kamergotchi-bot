package main

import (
	"log"
	"strconv"
	"time"
)

// ClaimLoop recusively keeps claiming rewards
func ClaimLoop(game Game, api *GameAPI) {
	resetTime, err := time.Parse(time.RFC3339, game.ClaimReset)
	if err != nil {
		handleClaimError(api, err)
		return
	}

	duration := calcClaimWaitDuration(resetTime)

	if duration > 0 {
		log.Println("Making next claim in " + duration.String() + "...")
		time.Sleep(duration)
	}

	game, err = api.ClaimReward()
	if err != nil {
		handleClaimError(api, err)
		return
	}

	log.Println("Claimed reward, new score: " + strconv.Itoa(game.Score) + ".")
	ClaimLoop(game, api)
}

func calcClaimWaitDuration(resetTime time.Time) time.Duration {
	timeToWait := resetTime.Sub(time.Now())
	return timeToWait + 5*time.Second + randomSeconds(3)
}

func handleClaimError(api *GameAPI, err error) {
	log.Println("Error in claim loop occurred, restarting in 15 seconds: " + err.Error())
	time.Sleep(15 * time.Second)

	game, err := api.GetGameInfo()
	if err != nil {
		handleClaimError(api, err)
	}

	ClaimLoop(game, api)
}
