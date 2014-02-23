package errors

import (
	"net/http"
)

func GetError(resp *http.Response, provider string) error {

	if resp.StatusCode/100 < 4 {
		return nil
	}


	if provider == "baidu" {
		err := new(BaiduError)
		err.Parse(resp.Body)
		defer resp.Body.Close()

		err.StatusCode = resp.StatusCode

		return err
	}

	return nil
}

