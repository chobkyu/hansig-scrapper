package router

import (
	"learngo/github.com/chobkyu/hansik/scrapper"
	"net/http"

	"github.com/labstack/echo"
)

var locate = [12]string{"서울", "인천", "김포", "대구", "세종", "부산", "경주", "광주", "대전", "성남", "전주", "울산"}

func GetDataAtGoogle(c echo.Context) error {
	for _, loc := range locate {
		scrapper.Scrap(loc)
	}
	return c.JSON(http.StatusOK, nil)
}
