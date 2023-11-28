package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var baseURL string = "https://pcmap.place.naver.com/place/list?query=%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98&x=127.59770646143932&y=35.845355799999766&clientX=126.942428&clientY=37.485309&bounds=125.88385325590485%3B33.001391604349834%3B129.39435450781872%3B38.526144869437104&ts=1700727250844&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98"

// define a custom data type for the scraped data
type Product struct {
	name, price string
}

func test() {
	// initialize a Chrome browser instance on port 4444
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

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	//https://pcmap.place.naver.com/place/list?query=%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98&x=127.49842900465427&y=35.66375922262287&clientX=126.942428&clientY=37.485309&bounds=125.38429339311597%3B32.13406719963322%3B129.7146967713403%3B38.96469205760175&ts=1700809713850&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98/
	// visit the target page
	err = driver.Get("https://pcmap.place.naver.com/restaurant/list?query=%ED%95%9C%EC%8B%9D%20%EB%B7%94%ED%8E%98&x=127.72748063120486&y=35.872664719463&clientX=126.942428&clientY=37.485309&bounds=122.42976628246777%3B32.77160273252248%3B129.5895708661264%3B38.786695685258934&ts=1701050131602&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%20%EB%B7%94%ED%8E%98")
	if err != nil {
		log.Fatal("Error:", err)
	}

	scrollingScript := `

	 // scroll down the page 10 times
	 const scrolls = 10
	 let scrollCount = 0

   // scroll down and then wait for 0.5s
	 const scrollInterval = setInterval(() => {
	  window.scrollTo(0, document.body.scrollHeight)
	  scrollCount++

	  if (scrollCount === scrolls) {
	  clearInterval(scrollInterval)
	  }
	 }, 500)
	 `
	_, err = driver.ExecuteScript(scrollingScript, []interface{}{})

	if err != nil {
		log.Fatal("Error:", err)
	}

	// wait up to 10 seconds for the 60th product to be on the page
	err = driver.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		lastProduct, _ := driver.FindElement(selenium.ByCSSSelector, "li")
		if lastProduct != nil {
			return lastProduct.IsDisplayed()
		}
		return false, nil
	}, 20*time.Second)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// retrieve the page raw HTML as a string
	// and logging it

	// html, err := driver.PageSource()
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// fmt.Println(html)

	// testEle, err := driver.FindElement(selenium.ByCSSSelector, ".blind")
	// if err != nil {
	// 	log.Fatal("Error : ", err)
	// }

	// fmt.Println(testEle.Text())

	// iframe, err := driver.FindElement(selenium.ByCSSSelector, "body > div > div > #app-layout > #section_content > div > .sc-1wsjitl > div > div ")
	// if err != nil {
	// 	log.Fatal("Error : ", err)
	// }

	//fmt.Println(iframe)
	//driver.SwitchFrame("#searchIframe")

	// html, err := driver.PageSource()
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// fmt.Println(html)

	productElements, err := driver.FindElements(selenium.ByCSSSelector, "li")
	if err != nil {
		log.Fatal("Error : ", err)
	}

	//fmt.Println(productElements)
	for _, productElement := range productElements {
		nameElement, err := productElement.FindElement(selenium.ByCSSSelector, ".TYaxT")

		name, err := nameElement.Text()

		if err != nil {
			log.Fatal("Error : ", err)
		}

		fmt.Println(name)
	}

}

func sel() {
	// where to store the scraped data
	var products []Product

	// initialize a Chrome browser instance on port 4444
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

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// visit the target page
	err = driver.Get("https://pcmap.place.naver.com/place/list?query=%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98&x=127.49842900465427&y=35.66375922262287&clientX=126.942428&clientY=37.485309&bounds=125.38429339311597%3B32.13406719963322%3B129.7146967713403%3B38.96469205760175&ts=1700809713850&mapUrl=https%3A%2F%2Fmap.naver.com%2Fp%2Fsearch%2F%ED%95%9C%EC%8B%9D%EB%B7%94%ED%8E%98")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// perform the scrolling interaction
	scrollingScript := `

	 // scroll down the page 10 times
	 const scrolls = 10
	 let scrollCount = 0

   // scroll down and then wait for 0.5s
	 const scrollInterval = setInterval(() => {
	  window.scrollTo(0, document.body.scrollHeight)
	  scrollCount++

	  if (scrollCount === scrolls) {
	  clearInterval(scrollInterval)
	  }
	 }, 500)
	 `
	_, err = driver.ExecuteScript(scrollingScript, []interface{}{})

	if err != nil {
		log.Fatal("Error:", err)
	}

	// wait up to 10 seconds for the 60th product to be on the page
	err = driver.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		lastProduct, _ := driver.FindElement(selenium.ByCSSSelector, ".UEzoS rTjJo cZnHG")
		if lastProduct != nil {
			return lastProduct.IsDisplayed()
		}
		return false, nil
	}, 20*time.Second)
	if err != nil {
		log.Fatal("Error:", err)
	}
	// select the product elements
	productElements, err := driver.FindElements(selenium.ByCSSSelector, ".UEzoS rTjJo cZnHG")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// iterate over the product elements
	// and extract data from them
	for _, productElement := range productElements {
		// select the name and price nodes
		nameElement, err := productElement.FindElement(selenium.ByCSSSelector, ".TYaxT")
		priceElement, err := productElement.FindElement(selenium.ByCSSSelector, ".Pb4bU")

		// extract the data of interest
		name, err := nameElement.Text()
		price, err := priceElement.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}

		// add the scraped data to the list
		product := Product{}
		product.name = name
		product.price = price
		products = append(products, product)
	}

	fmt.Println(products)

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

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)

	}
}
