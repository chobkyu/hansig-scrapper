package main

import (
	"learngo/github.com/chobkyu/hansik/config"
	"learngo/github.com/chobkyu/hansik/controller"
	"learngo/github.com/chobkyu/hansik/handlers"
	"learngo/github.com/chobkyu/hansik/router"

	"github.com/labstack/echo"
)

// https://velog.io/@kimdy0915/Selenium%EC%9C%BC%EB%A1%9C-%EB%84%A4%EC%9D%B4%EB%B2%84-%EC%A7%80%EB%8F%84-%ED%81%AC%EB%A1%A4%EB%A7%81%ED%95%98%EA%B8%B0
// https://www.zenrows.com/blog/selenium-golang#parse-the-data
func main() {
	e := echo.New()

	//GET
	e.GET("/", handlers.Home)
	e.GET("/getData", router.GetDataAtGoogle)
	e.GET("/ttt", handlers.Test)
	//POST
	e.POST("/test", handlers.TestDB)

	//init db connection
	//storage.InitDB()

	//connect to database gorm
	config.DatabaseInint()
	gorm := config.DB()

	dbGorm, err := gorm.DB()

	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	bookRoute := e.Group("/book")
	bookRoute.POST("/", controller.CreateBook)
	bookRoute.GET("/:id", controller.GetBook)
	bookRoute.PUT("/:id", controller.UpdateBook)
	bookRoute.DELETE("/:id", controller.DeleteBook)
	bookRoute.GET("/hansik", controller.CreateData)

	// hansikRoute := e.Group("/hansik")
	// hansikRoute.GET("/", controller.CreateData)

	e.Logger.Fatal(e.Start(":1323"))
}
