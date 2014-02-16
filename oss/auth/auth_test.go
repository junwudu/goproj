package auth

import (
	"testing"
	"fmt"
)


func TestSign(t *testing.T) {
	a := Auth{AccessInfo{"mzx6uUfGhzNiidxNuRjaEmTc", []byte("TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l")}, "bcs.duapp.com", "baidu"}
	p := SignParameter{"GET", "mytestoss", "/baiduyun/something.txt", "1393344000", "", "152444458"}

	get := "http://bcs.duapp.com/mytestoss/baiduyun%2Fsomething.txt?sign=MBOTS:mzx6uUfGhzNiidxNuRjaEmTc:r2YTDEVKbKVA6P1YcMk6cePsWxI%3D&time=1393344000&size=152444458"

	if url, err := a.SignedUrl(&p); url != get {
		fmt.Println(get)
		fmt.Println(url)
		t.Fatal(url, err)
	}

}
