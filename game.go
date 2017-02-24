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

// Gotchi contains basic info about the gotchi
type Gotchi struct {
	Name  string `json:"displayName"`
	Party string `json:"partyText"`
}

func (g Gotchi) getInfo() string {
	return g.Name + " from " + g.Party
}
