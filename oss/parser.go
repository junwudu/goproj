package oss

import (
	"io"
)

type Parser interface {
	Parse(reader io.Reader, result interface {}) error
}

