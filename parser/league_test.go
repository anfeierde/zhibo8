package parser

import (
	"io/ioutil"
	"testing"
)

func TestLeague(t *testing.T) {
	items := []string{"利物浦", "切尔西", "曼城"}
	urls := []string{
		"https://db.qiumibao.com/f/index/team?id=5",
		"https://db.qiumibao.com/f/index/team?id=138",
		"https://db.qiumibao.com/f/index/team?id=139",
	}

	content, err := ioutil.ReadFile("./league_data.json")
	if err != nil {
		panic(err)
	}

	result := League(content)

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
