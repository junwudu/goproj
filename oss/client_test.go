package oss

import (
	"testing"
)


var client = GetClient("baidu")

var parser BaiduParser

func TestListBucket(t *testing.T) {
	var aa Bucket
	aa.Name = "test18121"
	aa.Client = &client
	var bb Bucket
	bb.Name = "test18122"
	bb.Client = &client
	var cc Bucket
	cc.Name = "testbucket18123"
	cc.Client = &client

	aa.Create()
	bb.Create()
	cc.Create()

	defer aa.Delete()
	defer bb.Delete()
	defer cc.Delete()

	buckets, err := client.ListBucket(&parser)
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	for _, bucket := range buckets {
		if bucket.Name == aa.Name || bucket.Name == bb.Name || bucket.Name == cc.Name {
			count++
			continue
		}
	}

	if count < 3 {
		t.Fatal(buckets)
	}

}
