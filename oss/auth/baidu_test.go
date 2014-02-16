package provider

import (
	"testing"
	"github.com/junwudu/oss/auth"
	"fmt"
)



func TestBaiduDo(t *testing.T) {
	a := auth.Auth{auth.AccessInfo{"key", []byte("secret")}, "bcs.baidu.com", "baidu"}
	p := auth.SignParameter{"PUT", "/", "", "", "", ""}

	fmt.Println(BaiduDo(&a, &p))
}
