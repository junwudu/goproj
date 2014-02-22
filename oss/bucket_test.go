package oss

import (
	"testing"
	"fmt"
)

func TestListBucket(t *testing.T) {

	c := NewClient("baidu")

	var b BaiduParser

	buckets, err := ListBucket(&c, &b)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(buckets)
}
