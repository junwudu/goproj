package main

import (
	"time"
	"fmt"
	"github.com/lealife/spider"
)

func main() {
	start := time.Now()
	fmt.Println("start...")

	lea := spider.NewLeaSpider()
	lea.Fetch("http://www.lealife.com/", "E:/lealife")

	fmt.Printf("time cost %v\n", time.Now().Sub(start))
}
