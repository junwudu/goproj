package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
	"os"
	//"code.google.com/p/go.net/html"
	"github.com/PuerkitoBio/goquery"
)


func main() {
	res, err := http.Get("http://pp.163.com/wenyu8qiao/pp/11565003.html#35932037")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.m-picsetitem").Each(func(idx int, s *goquery.Selection) {
		src, exists := s.Find("div.pic-area img").Attr("data-lazyload-src")
		if exists {
			fmt.Printf("%d, %v\n", idx, src)
			pic, err := http.Get(src)
			defer pic.Body.Close()

			if err != nil {
				log.Fatal(err)
			}

			data, err := ioutil.ReadAll(pic.Body)

			if err != nil {
				log.Fatal(err)
			}

			fname := "E:\\" + strconv.FormatInt(int64(idx), 10) + ".jpg"
			fmt.Printf("%s, %d\n", fname, len(data))
			ioutil.WriteFile(fname, data, os.ModePerm)


		}

	})


}


