package oss

import (
	"github.com/junwudu/goproj/oss/auth"
	"fmt"
	"net/http"
	"github.com/junwudu/goproj/oss/errors"
	"time"
)


type Client struct {
	Provider auth.Provider
}


var GetClient = func() func(provider string) Client {
	var client Client

	return func(provider string) Client {

		if client.Provider != nil && client.Provider.Name() == provider {
			return client
		}

		client.Provider = auth.GetProvider(provider)
		return client
	}
}()


func (client Client) SignedUrl(method string, bucket string, object string, time string, ip string, size string) (url string, err error) {

	param := auth.SignParameter{method, bucket, object, time, ip, size}
	url, err = client.Provider.Sign(&param)
	return
}

func (client Client) ObjectUrl(object *Object) string {
	return fmt.Sprintf("http://%s/%s%s", client.Provider.Host(), object.Bucket.Name, object.Name)
}


func (client Client) NewBucket(name string) Bucket {
	var bucket Bucket
	bucket.Client = &client
	bucket.Name = name
	bucket.Born = time.Now()
	return bucket
}


func (client Client) ListBucket(parser Parser) (buckets []Bucket, err error) {

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
