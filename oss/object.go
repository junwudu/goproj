package oss

import (
	"time"
	"net/http"
	"github.com/junwudu/goproj/oss/errors"
	"fmt"
	"io"
)


type Object struct {
	/*object name starting with '/'*/
	Name string

	/*alias of Name, used as download name*/
	Alias string

	/*start with '/', without bucket part */
	ParentDir string

	IsDir bool

	/*modify time */
	ModifyTime time.Time

	/*bucket that in */
	Bucket string

	/*starting pos in this bucket*/
	Pos uint64

	/*content type (common this field in http header)*/
	Type string

	Data io.Reader

	/*Range start*/
	Start uint64

	/*size of object by byte*/
	Size uint64

	MD5 string

	/*user defined info*/
	Meta map[string]string

	/*Acl of this object*/
	Acl string
}


func (object *Object) ValidPut() (err error) {

	if len(object.Name) < 2 || object.Name[0] != '/' {
		err = errors.Error("object name is valid fail: " + object.Name)
	}

	if object.Bucket == "" {
		err = errors.Error("bukcet is not set")
	}

	if object.Data == nil {
		err = errors.Error("not data")
	}

	if object.Type == "" {
		err = errors.Error("content type is set")
	}
	return
}


func ListObject(client *Client, bucket string, parser Parser) (objects []Object, err error) {
	url, err := client.SignedUrl("GET", bucket, "/", "", "", "")

	fmt.Println(url)

	if err != nil {
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = errors.GetError(resp, client.Provider)

	if err == nil {
		err = parser.Parse(resp.Body, &objects)
	}

	return
}



func DeleteObject(client *Client, bucket string, object string) (err error) {
	url, err := client.SignedUrl("DELETE", bucket, object, "", "", "")
	fmt.Println(url)
	if err != nil {
		return
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if (err != nil) {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, client.Provider)

	return
}


func PutObject(client *Client, bucket string, object *Object) (err error) {

	err = object.ValidPut()
	if err != nil {
		return
	}

	url, err := client.SignedUrl("POST", bucket, object.Name, "", "", "")
	fmt.Println(url)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, object.Data)

	if err != nil {
		return
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", object.Type)
	}

	if req.Header.Get("Content-Disposition") == "" {
		req.Header.Set("Content-Disposition", object.Alias)
	}

	if req.Header.Get("x-bs-acl") == "" {
		req.Header.Set("x-bs-acl", object.Acl)
	}

	if len(object.Meta) > 0 {
		for k, v := range object.Meta {
			req.Header.Set(k, v)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, client.Provider)

	if err == nil {
		md5 := resp.Header.Get("Content-MD5")
		if md5 == "" {
			err = errors.Error("put error! content md5 empty")
		} else {
			object.MD5 = md5
		}
	}

	return
}


func GetObject(client *Client, bucket string, object *Object) {

}
