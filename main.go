package main

import (
	"learngo/github.com/chobkyu/hansik/scrapper"
)

// define a custom data type for the scraped data
type Product struct {
	name, price string
}

var locate = [12]string{"서울", "인천", "김포", "대구", "세종", "부산", "경주", "광주", "대전", "성남", "전주", "울산"}

// https://velog.io/@kimdy0915/Selenium%EC%9C%BC%EB%A1%9C-%EB%84%A4%EC%9D%B4%EB%B2%84-%EC%A7%80%EB%8F%84-%ED%81%AC%EB%A1%A4%EB%A7%81%ED%95%98%EA%B8%B0
// https://www.zenrows.com/blog/selenium-golang#parse-the-data
func main() {
	//c := make(chan int)
	for _, loc := range locate {
		scrapper.Scrap(loc)
	}
	//scrap()
}
