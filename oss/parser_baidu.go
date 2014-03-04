package oss

import (
	"io"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
	"net/http"
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


type baiduObject struct {
	Version_Key string
	Object string
	SuperFile string
	Size string
	Parent_dir string
	Is_dir string
	MDatetime string
	Ref_Key string
	Content_md5 string
}

type baiduObjects struct {
	Object_Total uint64
	Bucket string
	Start uint64
	Limit uint64
	Object_List []baiduObject
}


func (p *BaiduParser) Parse(reader io.Reader, result interface{}) error {

	switch result.(type) {
	default:

		return nil
	case *[]Bucket:
		return parseListBucket(reader, result.(*[]Bucket))

	case *[]Object:
		return parseListObject(reader, result.(*[]Object))
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

	bList := make([]Bucket, len(jObj))
	*buckets = bList

	for i, b := range jObj {
		bList[i].Name = b.Bucket_Name
		bList[i].Born, _ = time.Parse(http.TimeFormat, b.CDateTime)
		bList[i].Capacity, _ =  strconv.ParseUint(b.Total_Capacity, 10, 64)
		bList[i].Used, _ = strconv.ParseUint(b.Used_Capacity, 10, 64)
		bList[i].Status = b.Status
		bList[i].Location = b.Region
	}

	return nil
}


func parseListObject(reader io.Reader, objects *[]Object) error {
	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return err
	}

	var objList baiduObjects

	err = json.Unmarshal(data, &objList)
	if err != nil {
		return err
	}

	oList := make([]Object, len(objList.Object_List))
	*objects = oList

	for i, o := range objList.Object_List {
		var b Bucket
		b.Name = objList.Bucket
		oList[i].Bucket = b
		oList[i].Pos = objList.Start + uint64(i)
		oList[i].IsDir = o.Is_dir != "0"
		oList[i].Name = o.Object
		oList[i].ParentDir = o.Parent_dir

		t, err := strconv.ParseInt(o.MDatetime, 10, 64)
		if err == nil {
			oList[i].ModifyTime = time.Unix(t, 0)
		}

		size, err := strconv.ParseUint(o.Size, 10, 64)
		if err == nil {
			oList[i].Size = size
		}

		oList[i].MD5 = o.Content_md5
	}

	return nil
}

