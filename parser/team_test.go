package parser

import (
	"io/ioutil"
	"testing"
)

func TestTeam(t *testing.T) {
	items := []string{"索兰克", "菲尔米诺", "英斯"}
	urls := []string{
		"https://db.qiumibao.com/f/index/player?pid=24",
		"https://db.qiumibao.com/f/index/player?pid=25",
		"https://db.qiumibao.com/f/index/player?pid=26",
	}

	content, err := ioutil.ReadFile("./team_data.json")
	if err != nil {
		panic(err)
	}
	result := Team(content)

	for i, item := range items {
		if result.Items[i] != item {
			t.Errorf("error item not match, %s, %s \n", item, result.Items[i])
		}
	}

	for i, url := range urls {
		if result.Requests[i].Url != url {
			t.Errorf("error url not match, %s, %s \n", url, result.Requests[i].Url)
		}
	}
}
