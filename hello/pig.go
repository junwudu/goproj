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

}

func main() {

	fmt.Println(url.QueryEscape("$-_.+!*'(),"))

	fmt.Println("%" + strings.ToUpper(strconv.FormatInt(int64([]byte("-")[0]), 16)))
}

