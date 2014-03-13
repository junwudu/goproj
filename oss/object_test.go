package oss

import (
	"testing"
)

//var client = GetClient("baidu")

func createObject() Object {
	client := GetClient("baidu")
	b := client.NewBucket("mytestoss")
	obj := b.NewObject("otest.txt")
	obj.setData([]byte(`{
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1"}`))
	return obj
}

func isExistObject(object *Object) bool {
	return object.Head() == nil
}


func TestPutObject(t *testing.T) {
	object := createObject()

	if isExistObject(&object) {
		err := object.Delete()
		if err != nil {
			t.Fatal(err)
		}
	}

	if isExistObject(&object) {
		t.Fatal("can't delete object : " + object.Name)
	}

	if err := object.Put(); err != nil {
		t.Fatal(err)
	}

	if !isExistObject(&object) {
		t.Fatal("can't put object : " + object.Name)
	}
}


func TestDeleteObject(t *testing.T) {
	object := createObject()

	if !isExistObject(&object) {
		if err := object.Put(); err != nil {
			t.Fatal(err)
		}
	}

	if !isExistObject(&object) {
		t.Fatal("can't put object : " + object.Name)
	}

	err := object.Delete()
	if err != nil {
		t.Fatal(err)
	}

	if isExistObject(&object) {
		t.Fatal("can't delete object : " + object.Name)
	}
}


func TestCopyObject(t *testing.T) {
	object := createObject()

	if !isExistObject(&object) {
		if err := object.Put(); err != nil {
			t.Fatal(err)
		}
	}

	dtb := client.NewBucket(object.Bucket.Name)
	var dstObject = dtb.NewObject("test/mytest" + object.Alias)


	err := dstObject.Copy(&object, true)
	if err != nil {
		t.Fatal(err)
	}

	if !isExistObject(&dstObject) {
		t.Error(dstObject.Name + "is not copied!")
	}
}

func TestGetObject(t *testing.T) {
	object := createObject()

	if !isExistObject(&object) {
		if err := object.Put(); err != nil {
			t.Fatal(err)
		}
	}

	err := object.Get()
	if err != nil {
		t.Fatal(err)
	}

	object.Location = "E:/"
	err = object.Get()
	if err != nil {
		t.Fatal(err)
	}
}


func TestHeadObject(t *testing.T) {
	object := createObject()

	if !isExistObject(&object) {
		if err := object.Put(); err != nil {
			t.Fatal(err)
		}
	}

	if !isExistObject(&object) {
		t.Fatal("head faild")
	}
}


