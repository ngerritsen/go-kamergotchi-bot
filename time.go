package main

import (
	"math/rand"
	"time"
)

func randomSeconds(n int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(n)) * time.Second
}
