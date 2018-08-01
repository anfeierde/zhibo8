package model

type WordInfo struct {
	StatusMessage
	Data Continent `json:"data"`
}

type Continent struct {
	European European `json:"欧洲联赛"`
	Asia     Asia     `json:"亚洲联赛"`
}

type European struct {
	England []League `json:"英格兰"`
	Germany []League `json:"德国"`
	France  []League `json:"法国"`
	Italy   []League `json:"意大利"`
	Spain   []League `json:"西班牙"`
}

type Asia struct {
	China     []League `json:"中国"`
	Japan     []League `json:"日本"`
	Korea     []League `json:"韩国"`
	Australia []League `json:"澳大利亚"`
	Qatar     []League `json:"卡塔尔"`
}

type League struct {
	Id   string `json:"id"`
	Name string `json:"name_cn"`
}
