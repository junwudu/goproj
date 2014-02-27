package errors

import (
	"net/http"
)

type ErrorParser interface {
	Parse(resp *http.Response) error
}


type _error struct {
	Msg string
}

func (e *_error) Error() string {
	return e.Msg
}


func Error(msg string) error {
	var err _error
	err.Msg = msg
	return err
}
