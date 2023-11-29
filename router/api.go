package router

import (
	"fmt"
	"learngo/github.com/chobkyu/hansik/scrapper"
	"net/http"

	"github.com/labstack/echo"
)

var locate = []string{"서울", "인천", "김포", "대구", "세종", "부산", "경주", "광주", "대전", "성남", "전주", "울산"}

type hansikdang struct {
	name string
	addr string
	star string
}

func GetDataAtGoogle(c echo.Context) error {
	var hansikArr []hansikdang
	for _, loc := range locate {
		hansik := scrapper.Scrap(loc)
		fmt.Println(hansik)
		//hansikArr = append(hansikArr, hansik)
	}

	fmt.Println(hansikArr)
	return c.JSON(http.StatusOK, nil)
}
