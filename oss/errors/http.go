package errors

import (
	"net/http"
	"github.com/junwudu/goproj/oss/auth"
)

func GetError(resp *http.Response, provider auth.Provider) error {

	if resp.StatusCode/100 < 4 {
		return nil
	}

	if provider.Name() == "baidu" {
		err := new(BaiduError)
		err.Parse(resp.Body)
		defer resp.Body.Close()

		err.StatusCode = resp.StatusCode

		return err
	}

	return nil
}

