package main

// Game contains all info about the game's state
type Game struct {
	Score    int    `json:"score"`
	Health   int    `json:"health"`
	CareLeft int    `json:"careLeft"`
	Stats    Stats  `json:"current"`
	Gotchi   Gotchi `json:"gotchi"`
}

// Stats contain the stats of the gotchi
type Stats struct {
	Food      int `json:"food"`
	Attention int `json:"attention"`
	Knowledge int `json:"knowledge"`
}
