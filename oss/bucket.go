package oss

import (
	"net/http"
	"time"
	"strconv"
	"github.com/junwudu/goproj/oss/errors"
)


type Bucket struct {
	/*bucket name*/
	Name string

	/*bucket status*/
	Status string

	/*bucket created timestamp*/
	Born string

	/*bucket total capacity*/
	Capacity string

	/*bucket used capacity*/
	Used string

	/*bucket located geometry region*/
	Location string
}


func (bucket Bucket) BornTime() time.Time {
	t, err := strconv.ParseInt(bucket.Born, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(t, 0)
}



func ListBucket(client *Client, parser Parser) (buckets []Bucket, err error) {

	url, err := client.SignedUrl("GET", "", "/", "", "", "")
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
		err = parser.Parse(resp.Body, &buckets)
	}

	return
}


func CreateBucket(client *Client, name string) (err error) {
	url, err := client.SignedUrl("PUT", name, "/", "", "", "")
	if err != nil {
		return
	}

	req, err := http.NewRequest("PUT", url, nil)
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


func DeleteBucket(client *Client, name string) (err error) {
	url, err := client.SignedUrl("DELETE", name, "/", "", "", "")
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
