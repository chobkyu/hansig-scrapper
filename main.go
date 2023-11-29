package main

import (
	"learngo/github.com/chobkyu/hansik/handlers"
	"learngo/github.com/chobkyu/hansik/router"
	"learngo/github.com/chobkyu/hansik/storage"

	"github.com/labstack/echo"
)

// https://velog.io/@kimdy0915/Selenium%EC%9C%BC%EB%A1%9C-%EB%84%A4%EC%9D%B4%EB%B2%84-%EC%A7%80%EB%8F%84-%ED%81%AC%EB%A1%A4%EB%A7%81%ED%95%98%EA%B8%B0
// https://www.zenrows.com/blog/selenium-golang#parse-the-data
func main() {
	e := echo.New()

	//GET
	e.GET("/", handlers.Home)
	e.GET("/getData", router.GetDataAtGoogle)

	//POST

	//init db connection
	storage.InitDB()

	e.Logger.Fatal(e.Start(":1323"))
}
