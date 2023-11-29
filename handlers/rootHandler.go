package handlers

import "github.com/labstack/echo"

func Home(c echo.Context) error {
	return c.File("../html/home.html")

}
