package main

import "strconv"

import "time"

import "fmt"

import (

)


type S struct {
	A string
	B string
}

func main() {

	x := "1390308721"
	i, _ :=strconv.ParseInt(x, 10, 32)
	println(i)
	t := time.Unix(i, 0)

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(time.Now().Unix())

}

