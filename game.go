package main

// Game contains all info about the game's state
type Game struct {
	Score      int            `json:"score"`
	Health     int            `json:"health"`
	CareLeft   int            `json:"careLeft"`
	ClaimLimit int            `json:"claimLimitSeconds"`
	CareLimit  int            `json:"careLimitSeconds"`
	CareReset  string         `json:"careReset"`
	ClaimReset string         `json:"claimReset"`
	Stats      map[string]int `json:"current"`
	Gotchi     Gotchi         `json:"gotchi"`
}
