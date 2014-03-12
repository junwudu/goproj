package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
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

	var x []byte
	x = nil

	fmt.Println(len(x))

	ss := S{3,5}
	println(&ss.y)
	sss := ss

	vv := &ss

	println(&sss.x)
	println(&vv.y)



}
