package oss

import (
	"testing"
	"fmt"
)


func TestListObject(t *testing.T) {

	c := GetClient("baidu")

	var b BaiduParser

	objList, err := ListObject(&c, "mytestoss", &b)

	fmt.Println(objList, err)
}
