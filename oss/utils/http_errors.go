package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func GetError(resp *http.Response, provider string) error {

	if resp.StatusCode/100 < 4 {
		return nil
	}

	var httpError HttpError

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("err for read resp")
		return httpError
	}

	defer resp.Body.Close()

	fmt.Println(string(data))
	_ = provider

	return httpError
}


type HttpError struct {

}


func (httpError HttpError) Error() string {
	return ""
}




