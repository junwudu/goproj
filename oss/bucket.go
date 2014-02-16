package oss

import (
	"time"
	"net/http"
)


func ListBucket(client *Client, parser *Parser) (buckets []Bucket, err error) {

	url := client.SignedUrl("GET", "/", "", "", "", "")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	buckets, err := parser.Parse(resp.Body)

	return
}


type Bucket struct {
	/*bucket name*/
	Name string

	/*bucket created timestamp*/
	CreatedDateTime time.Time

	/*bucket total capacity*/
	TotalCapacity uint64

	/*bucket used capacity*/
	UsedCapacity uint64

	/*bucket located geometry region*/
	Region string
}


func (bucket *Bucket) Create(client *Client) (err error) {
	url := client.SignedUrl("PUT", bucket.Name, "", "", "", "")
	req, err := http.NewRequest("PUT", url, nil)
	if (err != nil) {
		panic(err)
	}

	_, err := http.DefaultClient.Do(req)

	return
}


func (bucket *Bucket) Delete(client *Client) (err error) {
	url := client.SignedUrl("DELETE", bucket.Name, "", "", "", "")
	req, err := http.NewRequest("DELETE", url, nil)
	if (err != nil) {
		panic(err)
	}

	_, err := http.DefaultClient.Do(req)

	return
}
