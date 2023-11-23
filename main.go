package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://map.naver.com/p/search/%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98?c=6.41,0,0,0,dh"

func main() {
	getTest()
}

func getTest() int {
	//req, rErr := http.NewRequest("GET", baseURL, nil)
	//checkErr(rErr)

	// 프록시로 호출
	//purl, err := url.Parse(baseURL)
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}
	// res, err := client.Do(req)
	// checkErr(err)
	// checkCode(res)

	// defer res.Body.Close()

	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// checkErr(err)

	// fmt.Println(doc)

	// doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Html())
	// })

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	//fmt.Println(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchLi := doc.Find(".TYaxT")

	searchLi.Each(func(i int, sikdang *goquery.Selection) {
		fmt.Println(sikdang.Text())
	})

	return 0

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)

	}
}
