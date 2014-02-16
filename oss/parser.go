package oss

import (
	"io"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Parser interface {
	Parse(reader io.Reader) (result interface {}, err error)
}


type BaiduParser struct {

}

func (parser *BaiduParser) ParseTo(reader io.Reader, result interface {}) error {
	var f interface {}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &f)

	switch result.(type) {
	case []Bucket:
		fmt.Println("-----------------")
	}

	return nil
}
