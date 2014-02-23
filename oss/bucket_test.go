package oss

import (
	"testing"
)


func ensureNotExist(bucket string) {
	c := GetClient("baidu")
	DeleteBucket(&c, bucket)
}


func ensureExist(bucket string) {
	c := GetClient("baidu")
	CreateBucket(&c, bucket)
}


func isExist(bucket string) bool {
	c := GetClient("baidu")
	var b BaiduParser

	buckets, err := ListBucket(&c, &b)
	if err != nil {
		panic(err)
	}

	for _, b := range buckets {
		if b.Name == bucket {
			return true
		}
	}

	return false
}


func TestCreateBucket(t *testing.T) {
	c := GetClient("baidu")

	testBucketName := "testjunmm"

	ensureNotExist(testBucketName)

	err := CreateBucket(&c, testBucketName)

	if err != nil {
		t.Error(err)
	}

	if !isExist(testBucketName) {
		t.Error("is not created")
	}

}

func TestDeleteBucket(t *testing.T) {
	c := GetClient("baidu")
	testBucketName := "testjunmm"

	ensureExist(testBucketName)

	err := DeleteBucket(&c, testBucketName)

	if err != nil {
		t.Error(err)
	}

	if isExist(testBucketName) {
		t.Fatal("is not deleted")
	}
}

func TestListBucket(t *testing.T) {
	c := GetClient("baidu")


	aa := "test18121"
	bb := "test18122"
	cc := "testbucket18123"

	ensureExist(aa)
	ensureExist(bb)
	ensureExist(cc)

	var b BaiduParser

	buckets, err := ListBucket(&c, &b)
	if err != nil {
		t.Fatal(err)
	}

	for i, bucket := range buckets {
		if bucket.Name == aa || bucket.Name == bb || bucket.Name == cc {
			continue
		}
		if i < 3 {
			t.Error(bucket.Name + " is not exist")
		}
	}

	ensureNotExist(aa)
	ensureNotExist(bb)
	ensureNotExist(cc)

}



