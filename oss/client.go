package oss

import (
	"github.com/junwudu/oss/auth"
	"net/http"
)



type Client struct {
	Auth auth.Auth
}


func NewClient(provider auth.Provider) Client {
	var client Client
	client.Auth = auth.NewAuth(provider)

	return client
}


func (client *Client) GetProvider() string {
	return string(client.Auth.Provider)
}


func (client *Client) SignedUrl(method string, bucket string, object string, time string, ip string, size string) string {
	client.Auth.Method = method
	client.Auth.Bucket = bucket
	client.Auth.Object = object
	client.Auth.Time = time
	client.Auth.Ip = ip
	client.Auth.Size = size
	return client.Auth.SignedUrl()
}

