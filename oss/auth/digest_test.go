package auth

import (
	"testing"
)



func TestAuthSign(t *testing.T) {
	info := NewInfo("baidu")
	auth := Auth{*info, "Get", "mytestoss", "/object", "", "", ""}

	s := auth.SignedUrl()

	if s == "" {
		t.Fatal(s)
	}

}
