package oss

import (
	"testing"
	"fmt"
	"os"
	"path/filepath"
)

//var client = GetClient("baidu")

var bucket = "mytestoss"

//var parser BaiduParser

var content = `{
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1"}`

var fn = "otest.txt"

func ensureFile() (path string) {

	f, err := os.Create(fn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.WriteString(content)

	path, err = filepath.Abs(f.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}


func ensureObject(object *Object) bool {
	path := ensureFile()
	defer object.SetDataFromFile(path)()

	err := PutObject(&client, object)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}


func isExist(object *Object) bool {
	err := HeadObject(&client, object)
	return err == nil
}


func TestPutObject(t *testing.T) {
	var object Object
	object.SetName(fn)
	object.SetBucket(bucket)

	if isExist(&object) {
		err := DeleteObject(&client, &object)
		if err != nil {
			t.Fatal(err)
		}
	}

	if !ensureObject(&object) {
		t.Fatal("failed to put object!")
	}
}


func TestDeleteObject(t *testing.T) {
	var object Object
	object.SetName(fn)
	object.SetBucket(bucket)

	if !isExist(&object) {
		if !ensureObject(&object) {
			t.Error("can't put objcect ")
		}
	}

	err := DeleteObject(&client, &object)
	if err != nil {
		t.Fatal(err)
	}
}


func TestCopyObject(t *testing.T) {
	var object Object
	object.SetName(fn)
	object.SetBucket(bucket)

	if !isExist(&object) {
		if !ensureObject(&object) {
			t.Error("can't put objcect ")
		}
	}

	var dstObject Object
	dstObject.SetName("test/mytest"+object.Alias)
	dstObject.SetBucket(bucket)

	err := CopyObject(&client, &dstObject, &object, true)
	if err != nil {
		t.Fatal(err)
	}

	if !isExist(&dstObject) {
		t.Error(dstObject.Name + "is not copied!")
	}
}

func TestGetObject(t *testing.T) {
	var object Object
	object.SetName(fn)
	object.SetBucket(bucket)

	if !isExist(&object) {
		if !ensureObject(&object) {
			t.Fatal("can't put objcect ")
		}
	}

	err := GetObject(&client, &object)
	if err != nil {
		t.Fatal(err)
	}

	object.Location = "E:/"
	err = GetObject(&client, &object)
	if err != nil {
		t.Fatal(err)
	}
}


func TestHeadObject(t *testing.T) {
	var object Object
	object.SetName(fn)
	object.SetBucket(bucket)

	if !isExist(&object) {
		if !ensureObject(&object) {
			t.Fatal("can't put objcect ")
		}
	}

	if !isExist(&object) {
		t.Fatal("head faild")
	}
}


func TestListObject(t *testing.T) {
	var object Object
	object.SetName(fn)
	object.SetBucket(bucket)

	if !isExist(&object) {
		if !ensureObject(&object) {
			t.Fatal("can't put objcect ")
		}
	}

	objList, err := ListObject(&client, object.Bucket, &parser)
	if err != nil {
		t.Fatal(err)
	}

	object.SetName(fn)

	var flag bool
	for _, obj := range objList {
		if obj.Name == object.Name {
			flag = true
		}
	}

	if !flag {
		t.Fatal(object.Name + " is not contained in object list")
	}
}

