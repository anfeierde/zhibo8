package parser

import (
	"encoding/json"
	"reflect"

	"cz1y.com/learngo/zhibo8/model"
)

const teamsUrl = "https://db.qiumibao.com/f/index/teams?name="

func Word(content []byte) model.ParserResult {
	result := model.ParserResult{}

	res := model.WordInfo{}
	err := json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}

	if res.Status != 1 {
		panic(res.Message)
	}

	leagues := []model.League{}

	v := reflect.ValueOf(res.Data)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					values := v.Field(i).Field(j).Interface().([]model.League)
					leagues = append(leagues, values...)
				}
				continue
			}
		}
	}

	for _, league := range leagues {
		result.Items = append(result.Items, league.Name)
		result.Requests = append(result.Requests, model.Request{
			Url:        teamsUrl + league.Name,
			ParserFunc: League,
		})
	}
	return result

}
