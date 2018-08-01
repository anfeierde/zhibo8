package parser

import (
	"io/ioutil"
	"testing"
)

func TestWord(t *testing.T) {
	urls := []string{
		"https://db.qiumibao.com/f/index/teams?name=英冠",
		"https://db.qiumibao.com/f/index/teams?name=英乙",
		"https://db.qiumibao.com/f/index/teams?name=英甲",
	}
	items := []string{
		"英冠", "英乙", "英甲",
	}

	text, err := ioutil.ReadFile("./word_data.json")
	if err != nil {
		panic(err)
	}
	result := Word(text)

	for i, item := range items {
		if result.Items[i] != item {
			t.Errorf("error item not match, %s, %s\n", item, result.Items[i])
		}
	}

	for i, url := range urls {
		if result.Requests[i].Url != url {
			t.Errorf("error url not match, %s, %s\n", url, result.Requests[i].Url)
		}
	}

}
