package utils

import (
	"strings"
)


func UrlPreEncode(urlStr string) string {

	urlStr = strings.Replace(urlStr, "_", "/", -1)
	urlStr = strings.Replace(urlStr, "-", "+", -1)

	return urlStr
}
