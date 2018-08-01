package parser

import (
	"encoding/json"

	"cz1y.com/learngo/zhibo8/model"
)

func League(content []byte) model.ParserResult {
	baseUrl := "https://db.qiumibao.com/f/index/team?id="
	result := model.ParserResult{}

	res := model.LeagueInfo{}
	err := json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}

	if res.Status != 1 {
		panic(res.Message)
	}

	for _, team := range res.Data.Teams {
		result.Items = append(result.Items, team.Name)
		result.Requests = append(result.Requests, model.Request{
			Url:        baseUrl + team.Id,
			ParserFunc: Team,
		})
	}

	return result
}
