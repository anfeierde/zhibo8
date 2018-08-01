package fetch

import "testing"

func TestFetch(t *testing.T) {
	var testUrl = "https://data.zhibo8.cc"
	doc, err := Fetch(testUrl)
	if err != nil {
		panic(err)
	}

}
