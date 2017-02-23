package main

import "log"

const playerToken = "aesthetickz"

func main() {
	log.Println("Go Kamergotchi bot started.")

	res, err := ApiRequest(playerToken)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
