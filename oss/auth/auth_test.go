package auth

import (
	"testing"
	"fmt"
)


func TestSign(t *testing.T) {
	a := Auth{AccessInfo{"key", []byte("secret")}, "bcs.baidu.com", "baidu"}
	p := SignParameter{"PUT", "/", "", "", "", ""}
	fmt.Println(a.Sign(&p))

}
