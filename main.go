package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://pcmap.place.naver.com/place/list?query=%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98&x=127.59770646143932&y=35.845355799999766&clientX=126.942428&clientY=37.485309&bounds=125.88385325590485%3B33.001391604349834%3B129.39435450781872%3B38.526144869437104&ts=1700727250844&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98"

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

	fmt.Println(doc)

	searchLi := doc.Find(".TYaxT")

	//fmt.Println(searchLi[0])

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
