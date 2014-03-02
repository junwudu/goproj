package oss

import (
	"github.com/junwudu/goproj/oss/auth"
	"fmt"
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
