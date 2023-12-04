package handlers

import (
	"fmt"
	"learngo/github.com/chobkyu/hansik/models"
	"learngo/github.com/chobkyu/hansik/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.File("../html/home.html")

}

func TestDB(c echo.Context) error {
	test := models.Test{}

	fmt.Println(c.FormValue("name"))
	c.Bind(&test)
	newTest, err := repositories.InsertData(test)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newTest)
}

func Test(c echo.Context) error {
	test, err := repositories.Test()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusCreated, test)

}
