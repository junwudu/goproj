package oss

import (
	"io"
	"io/ioutil"
	"encoding/json"
)

type BaiduParser struct {

}


type baiduBucket struct {
	Bucket_Name string
	Status string
	CDateTime string
	Used_Capacity string
	Total_Capacity string
	Region string
}


func (p *BaiduParser) Parse(reader io.Reader, result interface{}) error {

	switch result.(type) {
	default:

		return nil
	case *[]Bucket:
		return parseListBucket(reader, result.(*[]Bucket))

	}

}


func parseListBucket(reader io.Reader, buckets *[]Bucket) error {

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return err
	}

	var jObj []baiduBucket
	err = json.Unmarshal(data, &jObj)
	if err != nil {
		return err
	}

	bs := make([]Bucket, len(jObj))
	*buckets = bs

	for i, b := range jObj {
		bs[i].Name = b.Bucket_Name
		bs[i].Born = b.CDateTime
		bs[i].Capacity =  b.Total_Capacity
		bs[i].Used = b.Used_Capacity
		bs[i].Status = b.Status
		bs[i].Location = b.Region
	}

	return nil
}



