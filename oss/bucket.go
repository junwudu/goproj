package oss

import (
	"net/http"
	"time"
	"github.com/junwudu/goproj/oss/errors"
)


type Bucket struct {
	Client *Client
	/*bucket name*/
	Name string

	/*bucket status*/
	Status string

	/*bucket created timestamp*/
	Born time.Time

	/*bucket total capacity*/
	Capacity uint64

	/*bucket used capacity*/
	Used uint64

	/*bucket located geometry region*/
	Location string

	/*bucket access level*/
	Acl string

	Objects []Object
}

// create object from it's name,
// bucket within and the content.
// input maybe nil
func (bucket *Bucket) NewObject(name string) Object {
	var object Object
	object.Client = bucket.Client
	object.setName(name)
	object.Bucket = bucket
	return object
}



func (bucket *Bucket) Create() (err error) {
	url, err := bucket.Client.SignedUrl("PUT", bucket.Name, "/", "", "", "")
	if err != nil {
		return
	}

	req, err := http.NewRequest("PUT", url, nil)
	if (err != nil) {
		return
	}

	if bucket.Acl != "" {
		req.Header.Set(bucket.Client.Provider.Acl(), bucket.Acl)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, bucket.Client.Provider)

	return
}


func (bucket *Bucket) Delete() (err error) {
	url, err := bucket.Client.SignedUrl("DELETE", bucket.Name, "/", "", "", "")
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

	err = errors.GetError(resp, bucket.Client.Provider)

	return
}


func (bucket *Bucket) List(parser Parser) (err error) {
	url, err := bucket.Client.SignedUrl("GET", bucket.Name, "/", "", "", "")

	if err != nil {
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = errors.GetError(resp, bucket.Client.Provider)

	if err == nil {
		err = parser.Parse(resp.Body, &bucket.Objects)
	}
	return
}
