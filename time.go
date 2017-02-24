package main

import (
	"math/rand"
	"time"
)

// GetCareWaitDuration generates the duration to wait for the next care round
func GetCareWaitDuration(reset time.Time) time.Duration {
	timeToWait := reset.Sub(time.Now())
	return timeToWait + randomSeconds(60)
}

// GetClaimWaitDuration generates the duration to wait for the next claim
func GetClaimWaitDuration(reset time.Time) time.Duration {
	timeToWait := reset.Sub(time.Now())
	return timeToWait + 80*time.Second + randomSeconds(60)
}

func randomSeconds(n int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(n)) * time.Second
}
