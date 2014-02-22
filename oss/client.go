package oss

import (
	"github.com/junwudu/goproj/oss/auth"
)


type Client struct {
	authorize auth.Authorize
	Provider string
}


var GetClient = func() func(provider string) Client {
	var client Client

	return func(provider string) Client {

		if client.Provider == provider {
			return client
		}

		client.Provider = provider
		client.authorize = auth.Provide(provider)
		return client
	}
}()


func (client Client) SignedUrl(method string, bucket string, object string, time string, ip string, size string) (url string, err error) {
	param := auth.SignParameter{method, bucket, object, time, ip, size}
	url, err = client.authorize.Sign(&param)
	return
}

