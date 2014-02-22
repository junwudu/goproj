package oss

import (
	"github.com/junwudu/goproj/oss/auth"
)


type Client struct {
	authorize auth.Auth
}


func NewClient(provider string) Client {
	var client Client
	client.authorize = auth.Provide(provider)
	return client
}


func (client *Client) SignedUrl(method string, bucket string, object string, time string, ip string, size string) string {

	p := auth.SignParameter{method, bucket, object, time, ip, size}
	r, err := client.authorize.SignedUrl(&p)
	if err != nil {
		panic(err)
	}

	return r
}

