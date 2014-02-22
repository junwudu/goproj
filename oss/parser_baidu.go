package oss

import (
	"io"
	"io/ioutil"
	"encoding/json"
)

type BaiduParser struct {

}


func (p *BaiduParser) Parse(reader io.Reader, result interface{}) error {

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	return json.Unmarshal(data, result)
}



