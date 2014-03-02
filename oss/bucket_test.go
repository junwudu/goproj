package oss

import (
	"testing"
)


func isExistBucket(bucket string) bool {

	buckets, err := ListBucket(&client, &parser)
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

	testBucketName := "testjunmm"

	if isExistBucket(testBucketName) {
		err := DeleteBucket(&client, testBucketName)
		if err != nil {
			t.Fatal(err)
		}
	}

	err := CreateBucket(&client, testBucketName)

	if err != nil {
		t.Error(err)
	}

	if isExistBucket(testBucketName) {
		t.Skip(testBucketName + " is Created!")
	}

}


func TestDeleteBucket(t *testing.T) {
	testBucketName := "testjunmm"

	if !isExistBucket(testBucketName) {
		err := CreateBucket(&client, testBucketName)
		if err != nil {
			t.Fatal(err)
		}
	}

	err := DeleteBucket(&client, testBucketName)
	if err != nil {
		t.Error(err)
	}

	if !isExistBucket(testBucketName) {
		t.Skip(testBucketName + " is Deleted!")
	}
}


func TestListBucket(t *testing.T) {
	aa := "test18121"
	bb := "test18122"
	cc := "testbucket18123"

	CreateBucket(&client, aa)
	CreateBucket(&client, bb)
	CreateBucket(&client, cc)

	defer DeleteBucket(&client, aa)
	defer DeleteBucket(&client, bb)
	defer DeleteBucket(&client, cc)

	buckets, err := ListBucket(&client, &parser)
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

}



