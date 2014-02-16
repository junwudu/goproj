package main

import (
	"fmt"
	"reflect"
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

}

