package oss

import (
	"testing"
	"fmt"
	"github.com/junwudu/goproj/oss/errors"
)


func isExistBucket(bucket *Bucket) bool {

	buckets, err := client.ListBucket(&parser)
	if err != nil {
		panic(err)
	}

	for _, b := range buckets {
		if b.Name == bucket.Name {
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
	testBucket := client.NewBucket("testjunmm")

	if isExistBucket(&testBucket) {
		err := testBucket.Delete()
		if err != nil {
			t.Fatal(err)
		}
	}

	if isExistBucket(&testBucket) {
		t.Fatal(fmt.Sprintf("bucket %s has existed!", testBucket.Name))
	}

	err := testBucket.Create()
	if err != nil {
		t.Fatal(err)
	}

	if !isExistBucket(&testBucket) {
		t.Fatal("create bucket " + testBucket.Name + " is failed!")
	}

	//test recreate
	err = testBucket.Create()
	if err != nil {
		assertError(t, err, -1001, 403)
	}
}


func TestDeleteBucket(t *testing.T) {
	testBucket := client.NewBucket("testjunmm")

	if !isExistBucket(&testBucket) {
		err := testBucket.Create()
		if err != nil {
			t.Fatal(err)
		}
	}

	if !isExistBucket(&testBucket) {
		t.Fatal("bucket can not be created : " + testBucket.Name)
	}

	//put object
	b := client.NewBucket(testBucket.Name)
	obj := b.NewObject("otest.txt")
	obj.setData([]byte(`{
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1"}`))

	if err := obj.Put(); err != nil {
		t.Error("can't put object")
	} else {
		err := testBucket.Delete()
		if err != nil {
			assertError(t, err, -1007, 403)
		}

		err = obj.Delete()
		if err != nil {
			t.Fatal(err)
		}
	}

	err := testBucket.Delete()
	if err != nil {
		t.Fatal(err)
	}

	if isExistBucket(&testBucket) {
		t.Fatal("bucket can not be deleted : " + testBucket.Name)
	}

	err = testBucket.Delete()
	if err != nil {
		assertError(t, err, -42, 403)
	}
}



func TestListObject(t *testing.T) {
	object := createObject()

	if !isExistObject(&object) {
		if err := object.Put(); err != nil {
			t.Fatal(err)
		}
	}

	err := object.Bucket.List(&parser)
	if err != nil {
		t.Fatal(err)
	}


	var flag bool
	for _, obj := range object.Bucket.Objects {
		if obj.Name == object.Name {
			flag = true
		}
	}

	if !flag {
		t.Fatal(object.Name + " is not contained in object list")
	}
}



