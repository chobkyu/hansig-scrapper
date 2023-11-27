package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// define a custom data type for the scraped data
type Product struct {
	name, price string
}

// https://velog.io/@kimdy0915/Selenium%EC%9C%BC%EB%A1%9C-%EB%84%A4%EC%9D%B4%EB%B2%84-%EC%A7%80%EB%8F%84-%ED%81%AC%EB%A1%A4%EB%A7%81%ED%95%98%EA%B8%B0
// https://www.zenrows.com/blog/selenium-golang#parse-the-data
func main() {
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	//https://pcmap.place.naver.com/place/list?query=%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98&x=127.49842900465427&y=35.66375922262287&clientX=126.942428&clientY=37.485309&bounds=125.38429339311597%3B32.13406719963322%3B129.7146967713403%3B38.96469205760175&ts=1700809713850&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98/
	// visit the target page
	err = driver.Get("https://www.google.com/search?q=서울+한식+뷔페&sca_esv=585526170&sz=7&cs=0&biw=1327&bih=963&tbm=lcl&ei=jUZkZamyPPLf2roP0sCX0AI&ved=0ahUKEwjprNCs1eOCAxXyr1YBHVLgBSoQ4dUDCAk&oq=%EC%84%9C%EC%9A%B8+%ED%95%9C%EC%8B%9D+%EB%B7%94%ED%8E%98&gs_lp=Eg1nd3Mtd2l6LWxvY2FsIhTshJzsmrgg7ZWc7IudIOu3lO2OmEgAUABYAHAAeACQAQCYAQCgAQCqAQC4AQzIAQA&sclient=gws-wiz-local#rlfi=hd:;si:;mv:[[37.581255999999996,127.05252650000001],[37.4620829,126.83985559999999]];tbs:lrf:!1m4!1u3!2m2!3m1!1e1!1m4!1u2!2m2!2m1!1e1!2m1!1e2!2m1!1e3!3sIAE,lf:1,lf_ui:9")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// create a new remote client with the specified options
	// driver, err := selenium.NewRemote(caps, "")
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// html, err := driver.PageSource()
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	//fmt.Println(html)

	productElements, err := driver.FindElements(selenium.ByCSSSelector, ".dbg0pd")
	if err != nil {
		log.Fatal("Error : ", err)
	}

	//fmt.Println(productElements)
	for _, productElement := range productElements {
		nameElement, err := productElement.FindElement(selenium.ByCSSSelector, ".OSrXXb")

		name, err := nameElement.Text()

		if err != nil {
			log.Fatal("Error : ", err)
		}

		fmt.Println(name)
	}
}
