package errors

import (
	"net/http"
)

type ErrorParser interface {
	Parse(resp *http.Response) error
}
