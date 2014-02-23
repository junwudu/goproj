package main

import (
	"fmt"
	"net/url"
	"strconv"
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



func main() {

	fmt.Println(url.QueryEscape("$-_.+!*'(),"))

	fmt.Println("%" + strings.ToUpper(strconv.FormatInt(int64([]byte("-")[0]), 16)))

	var k []S

	k = []S {
		{34, 44},
		{546, 343},
	}

	fmt.Println(k)

	var m map[string]int
	m = make(map[string]int, 10)
	m["ss"] = 34332

	fmt.Println(strconv.FormatInt(34, 10))

	var bbb bool

	bbb = "0" != "0"

	fmt.Println(bbb)
}

