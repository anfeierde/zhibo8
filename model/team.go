package model

type TeamInfo struct {
	Players []Player `json:"players"`
}

type Player struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
