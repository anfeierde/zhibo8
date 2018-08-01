package model

type LeagueInfo struct {
	StatusMessage
	Data Teams `json:"data"`
}

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
