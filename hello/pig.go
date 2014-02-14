package main

import (
	"fmt"
	"runtime"
	"strconv"
//	"math/rand"
//	"math"
//	"reflect"
	"reflect"
	"strings"
)

func mem(suffix string) uint64 {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Printf("%s--> %v\n",suffix, mem.TotalAlloc)
	return mem.Alloc
}

var (
	aa = 33
	bb = 44
	cc = true
)

type A float64

func (a A) add(b A) (r A) {
	r = a + b
	return
}

func main() {
	s := mem("start")


	x := [][]byte {{3,4}, {5,6}}

	fmt.Printf("%v %T\n", x, x)

	z := make([][]byte, 3)

	fmt.Println(reflect.TypeOf(x) == reflect.TypeOf(z))

	x[0] = make([]byte, 2)
	fmt.Println(len(x), x)

	w := "i am a studio!"
	si := strings.Split(w, " ")
	fmt.Printf("%v\n", len(si))

	mp := make(map[string]int, 4)
	fmt.Println(mp["sfewfefew"])

	e := mem("end")

	println("diff -> " + strconv.FormatUint(e-s, 10))
}

