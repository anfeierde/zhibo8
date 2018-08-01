package parser

import (
	"encoding/json"

	"cz1y.com/learngo/zhibo8/model"
)

func Team(content []byte) model.ParserResult {
	baseUrl := "https://db.qiumibao.com/f/index/player?pid="

	result := model.ParserResult{}

	res := model.TeamInfo{}
	err := json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}

	for _, player := range res.Players {
		result.Items = append(result.Items, player.Name)
		result.Requests = append(result.Requests, model.Request{
			Url:        baseUrl + player.Id,
			ParserFunc: Player,
		})
	}

	return result
}
