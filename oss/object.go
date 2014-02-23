package oss

import (
	"time"
	"net/http"
	"github.com/junwudu/goproj/oss/errors"
	"fmt"
)


type Object struct {
	/*object name starting with '/'*/
	Name string

	Size uint64

	/*start with '/', without bucket part */
	ParentDir string

	IsDir bool

	/*modify time */
	ModifyTime time.Time

	MD5 string

	/*bucket that in */
	Bucket string

	/*starting pos in this bucket*/
	Pos uint64
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

func PutObject(client *Client, bucket string, object string) (err error) {

}
