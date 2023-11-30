package scrapper

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Hansikdang struct {
	Name string
	Addr string
	Star string
}

func Scrap(loc string) []Hansikdang {
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)

	checkErr(err)

	defer service.Stop()

	// configure the browser options

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
	}})

	driver, err := selenium.NewRemote(caps, "")
	checkErr(err)
	//https://pcmap.place.naver.com/place/list?query=%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98&x=127.49842900465427&y=35.66375922262287&clientX=126.942428&clientY=37.485309&bounds=125.38429339311597%3B32.13406719963322%3B129.7146967713403%3B38.96469205760175&ts=1700809713850&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98/
	// visit the target page
	err = driver.Get("https://www.google.com/search?q=" + loc + "+한식+뷔페&sca_esv=585526170&sz=7&cs=0&biw=1327&bih=963&tbm=lcl&ei=jUZkZamyPPLf2roP0sCX0AI&ved=0ahUKEwjprNCs1eOCAxXyr1YBHVLgBSoQ4dUDCAk&oq=%EC%84%9C%EC%9A%B8+%ED%95%9C%EC%8B%9D+%EB%B7%94%ED%8E%98&gs_lp=Eg1nd3Mtd2l6LWxvY2FsIhTshJzsmrgg7ZWc7IudIOu3lO2OmEgAUABYAHAAeACQAQCYAQCgAQCqAQC4AQzIAQA&sclient=gws-wiz-local#rlfi=hd:;si:;mv:[[37.581255999999996,127.05252650000001],[37.4620829,126.83985559999999]];tbs:lrf:!1m4!1u3!2m2!3m1!1e1!1m4!1u2!2m2!2m1!1e1!2m1!1e2!2m1!1e3!3sIAE,lf:1,lf_ui:9")
	checkErr(err)

	page := getPages(driver)

	hansikData := getDatas(page, loc)

	fmt.Println(hansikData)
	return hansikData

}

func getDatas(page int, loc string) []Hansikdang {

	var hansikTemp []Hansikdang

	for i := 1; i <= page; i++ {
		fmt.Println("--------------------------------------")
		fmt.Println("This page is data of ", i)

		caps := selenium.Capabilities{}
		caps.AddChrome(chrome.Capabilities{Args: []string{
			"--headless", // comment out this line for testing
		}})

		driver, err := selenium.NewRemote(caps, "")
		checkErr(err)

		if i == 1 {
			driver.Get("https://www.google.com/search?q=" + loc + "+한식+뷔페&sca_esv=585526170&sz=7&cs=0&biw=1015&bih=963&tbm=lcl&ei=qEllZc2PIcro-Aa8i5PYDQ&ved=0ahUKEwjNg8C5zOWCAxVKNN4KHbzFBNsQ4dUDCAk&oq=" + loc + "+한식+뷔페&gs_lp=Eg1nd3Mtd2l6LWxvY2FsIhTshJzsmrgg7ZWc7IudIOu3lO2OmEgAUABYAHAAeACQAQCYAQCgAQCqAQC4AQzIAQA&sclient=gws-wiz-local#rlfi=hd:;si:;mv:[[37.581255999999996,127.05252650000001],[37.4620829,126.83985559999999]];tbs:lrf:!1m4!1u3!2m2!3m1!1e1!1m4!1u2!2m2!2m1!1e1!2m1!1e2!2m1!1e3!3sIAE,lf:1,lf_ui:9")
			checkErr(err)

		} else {
			pageNum := strconv.Itoa((i - 1) * 20)
			driver.Get("https://www.google.com/search?q=" + loc + "+한식+뷔페&sca_esv=585526170&sz=7&cs=0&biw=1015&bih=963&tbm=lcl&ei=qEllZc2PIcro-Aa8i5PYDQ&ved=0ahUKEwjNg8C5zOWCAxVKNN4KHbzFBNsQ4dUDCAk&oq=" + loc + "+한식+뷔페&gs_lp=Eg1nd3Mtd2l6LWxvY2FsIhTshJzsmrgg7ZWc7IudIOu3lO2OmEgAUABYAHAAeACQAQCYAQCgAQCqAQC4AQzIAQA&sclient=gws-wiz-local#rlfi=hd:;si:;mv:[[37.656864899999995,127.13864459999999],[37.4778707,126.8077316]];start:20;start:" + pageNum)
			checkErr(err)

		}

		productElements, err := driver.FindElements(selenium.ByCSSSelector, ".rllt__details")
		if err != nil {
			fmt.Println("can't find details")
			fmt.Println(err)
			i++
		}

		//fmt.Println(productElements)
		for _, productElement := range productElements {
			fmt.Println(i)

			nameElement, err := productElement.FindElement(selenium.ByCSSSelector, ".OSrXXb")
			if err != nil {
				fmt.Println("error at ", i, " page")
				fmt.Println(err)
				break
			}
			name, err := nameElement.Text()
			fmt.Println(name)
			if err != nil {
				name = "no data"
				break
			}

			//fmt.Println("addr 앞")

			addrElement, err := productElement.FindElements(selenium.ByCSSSelector, "div")
			if err != nil {
				fmt.Println("no addr")
			}
			addrTemp := "주소 없음"
			if len(addrElement) > 2 {
				addrTemp, err = addrElement[2].Text()
				if err != nil {
					addrTemp = "주소 없음"
				}
			} else {
				fmt.Println("no addr")
			}
			addr := addrTemp

			//fmt.Println("star 앞")

			starElement, err := productElement.FindElement(selenium.ByCSSSelector, ".z3HNkc")
			var star string
			if err != nil {
				star = "리뷰 없음"
			} else {
				starcheck, err := starElement.GetAttribute("aria-label")
				if err == nil {
					star = starcheck
				}
			}

			// fmt.Println(name)
			// fmt.Println(addr)
			// fmt.Println(star)
			fmt.Println()
			var temp = Hansikdang{
				Name: name,
				Addr: addr,
				Star: star,
			}
			hansikTemp = append(hansikTemp, temp)
			time.Sleep(time.Second * 1)

		}

		time.Sleep(time.Second * 2)
	}
	return hansikTemp
}

func getPages(driver selenium.WebDriver) int {
	pages, err := driver.FindElements(selenium.ByCSSSelector, "td")
	checkErr(err)

	fmt.Println(len(pages))
	fmt.Println("Page!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	return len(pages)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
