package parser

import (
	"encoding/json"
	"fmt"

	"cz1y.com/learngo/zhibo8/model"
)

var count = 0
var first = true

func Player(content []byte) model.ParserResult {
	result := model.ParserResult{}

	res := model.PlayerInfo{}
	err := json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}

	//	result.Items = append(result.Items, res)
	fmt.Printf("detail: #%d, %v \n", count, res)

	if first {
		first = false
	}
	count++

	return result
}
