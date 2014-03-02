package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func M(p *[]string) {
	*p = make([]string, 2)
	*p = append(*p, "yes")
}


type S struct {
	x int
	y int
}


type A struct {
	b string
}

func main() {


	a := A{"3424"}

	fmt.Println(a)

	data := md5.Sum([]byte("nihao"))
	fmt.Println(hex.EncodeToString(data[0:]))

	x := "/baidu/nihao"

	fmt.Println(x[strings.LastIndex(x, "/") + 1  : len(x)])


}

