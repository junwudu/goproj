package oss

import (
	"testing"
	"fmt"
	"github.com/junwudu/goproj/oss/errors"
)


func isExistBucket(bucket string) bool {

	buckets, err := client.ListBucket(&parser)
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


func assertError(t *testing.T, err error, args...int) {
	switch err.(type) {
	default:
		t.Fatal(err)
	case *errors.BaiduError:
		e := err.(*errors.BaiduError)

		t.Log(e.Description)

		if e.ErrorCode != args[0] {
			t.Error(fmt.Sprintf("errorcode should be %d, but %d", -1001, e.ErrorCode))
		}

		if e.StatusCode != args[1] {
			t.Error(fmt.Sprintf("err statuscode should be %d, but %d", 403, e.StatusCode))
		}
	}
}


func TestCreateBucket(t *testing.T) {
	var testBucket Bucket
	testBucket.Name = "testjunmm"
	testBucket.Client = &client

	if isExistBucket(testBucket.Name) {
		err := testBucket.Delete()
		if err != nil {
			t.Fatal(err)
		}
	}

	if isExistBucket(testBucket.Name) {
		t.Fatal(fmt.Sprintf("bucket %s has existed!", testBucket.Name))
	}

	err := testBucket.Create()
	if err != nil {
		t.Fatal(err)
	}

	if !isExistBucket(testBucket.Name) {
		t.Fatal("create bucket " + testBucket.Name + " is failed!")
	}

	//test recreate
	err = testBucket.Create()
	if err != nil {
		assertError(t, err, -1001, 403)
	}
}


func TestDeleteBucket(t *testing.T) {
	var testBucket Bucket
	testBucket.Name = "testjunmm"
	testBucket.Client = &client

	if !isExistBucket(testBucket.Name) {
		err := testBucket.Create()
		if err != nil {
			t.Fatal(err)
		}
	}

	if !isExistBucket(testBucket.Name) {
		t.Fatal("bucket can not be created : " + testBucket.Name)
	}

	//put object
	var object Object
	object.SetName(fn)
	object.SetBucket(testBucket.Name)
	if !ensureObject(&object) {
		t.Error("can't put object")
	} else {
		err := testBucket.Delete()
		if err != nil {
			assertError(t, err, -1007, 403)
		}

		err = DeleteObject(&client, &object)
		if err != nil {
			t.Fatal(err)
		}
	}

	err := testBucket.Delete()
	if err != nil {
		t.Fatal(err)
	}

	if isExistBucket(testBucket.Name) {
		t.Fatal("bucket can not be deleted : " + testBucket.Name)
	}

	err = testBucket.Delete()
	if err != nil {
		assertError(t, err, -42, 403)
	}
}


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



