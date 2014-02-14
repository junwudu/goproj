package oss

import (
	"testing"
)


func TestListBuckets(t *testing.T) {
	client := NewClient("baidu")
	client.ListBuckets(BaiduParser{})
}
