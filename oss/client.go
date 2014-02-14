package oss

import (
	"github.com/junwudu/oss/auth"
	"net/http"
)



type Client struct {
	Auth auth.Auth
}

type Bucket string

type Object struct {
	Bucket Bucket
	Object string
}


func NewClient(provider auth.Provider) Client {
	var client Client
	client.Auth = auth.NewAuth(provider)

	return client
}


func (client *Client) GetProvider() string {
	return string(client.Auth.Provider)
}


func (client *Client) Sign(method string, bucket string, object string, time string, ip string, size string) string {
	client.Auth.Method = method
	client.Auth.Bucket = bucket
	client.Auth.Object = object
	client.Auth.Time = time
	client.Auth.Ip = ip
	client.Auth.Size = size
	return client.Auth.SignedUrl()
}


func (client *Client) ListBuckets(parser Parser) (buckets []Bucket) {
	url := client.Sign("GET", "", "/", "", "", "")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	err := parser.Parse(resp.Body, buckets)
	if err != nil {
		panic(err)
	}

	return []Bucket(buckets)

}


