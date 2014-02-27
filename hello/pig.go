package main

import (
	"fmt"
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
	m := make(map[string]string, 3)
	m["33"] = "3423"

	fmt.Println(len(m))
}

