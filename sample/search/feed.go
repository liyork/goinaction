package search

import (
	"os"
)

const dataFile = "data/data.json"

//结构
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//参数(),返回值[]*Feed, error
//*指针，当前和调用方公用指针所指内容
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	//todo
	//err = json.NewDecoder(file, json.Decoder{&feeds})

	return feeds, err
}
