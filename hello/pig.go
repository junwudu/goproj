package main

import (
	"fmt"
	"reflect"
	"strings"
	"github.com/junwudu/oss/utils"
)


func G() (result interface {}) {

	m := map[string]string{"a": "pp", "b": "cc"}

	result = interface {}(m)

	return
}


type S struct {
	x int
	y int
}


func (s *S) setX(x int) {
	s.x = x
}


func main() {
	fmt.Printf("%T-->%v", G(), reflect.TypeOf(G()))

	var s S
	s.setX(34)
	fmt.Println(s)

	y := "yes"
	fmt.Println(y[0])

	fmt.Println(strings.TrimLeft("/", "/"))

	fmt.Println(utils.Unescape("XhAUN1n6ilIR57Q9xEu3%2Bpvn5Uo%3D"))
}

