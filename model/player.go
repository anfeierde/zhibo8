package model

type PlayerInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	NameEn      string `json:"name_en"`
	Number      string `json:"number"`
	Position    string `json:"position"`
	Nationality string `json:"Nationality"`
	Age         string `json:"age"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Team        string `json:"team_name"`
	Goal        string `json:"goal"`
	Assist      string `json:"assist"`
	YellowCard  string `json:"yellow_card"`
	RedCard     string `json:"red_card"`
	BirtyDay    string `json:"birthday"`
	Foot        string `json:"habit_foot"`
}
