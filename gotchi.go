package main

// Gotchi contains basic info about the gotchi
type Gotchi struct {
	Name  string `json:"displayName"`
	Party string `json:"partyText"`
}

func (g Gotchi) getInfo() string {
	return g.Name + " from " + g.Party
}
