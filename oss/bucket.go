package oss

import (
	"net/http"
	"fmt"
	"time"
	"strconv"
)


func ListBucket(client *Client, parser Parser) (buckets []Bucket, err error) {

	url := client.SignedUrl("GET", "", "/", "", "", "")
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	err = parser.Parse(resp.Body, &buckets)

	return
}


type Bucket struct {
	/*bucket name*/
	Bucket_Name string

	Status string

	/*bucket created timestamp*/
	CdateTime string

	/*bucket used capacity*/
	Used_Capacity string

	/*bucket total capacity*/
	Total_Capacity string

	/*bucket located geometry region*/
	Region string
}


func (bucket *Bucket) Name() string {
	return bucket.Bucket_Name
}


/* Get Bucket Created datetime stamp */
func (bucket *Bucket) CreatedDateTime() time.Time {
	unix, _ := strconv.ParseInt(bucket.CdateTime, 10, 32)
	return time.Unix(unix, 0)
}



func (bucket *Bucket) Create(client *Client) (err error) {
	url := client.SignedUrl("PUT", bucket.Name(), "", "", "", "")
	req, err := http.NewRequest("PUT", url, nil)
	if (err != nil) {
		panic(err)
	}

	_, err = http.DefaultClient.Do(req)

	return
}


func (bucket *Bucket) Delete(client *Client) (err error) {
	url := client.SignedUrl("DELETE", bucket.Name(), "", "", "", "")
	req, err := http.NewRequest("DELETE", url, nil)
	if (err != nil) {
		panic(err)
	}

	_, err = http.DefaultClient.Do(req)

	return
}
