package router

import (
	"database/sql"
	"fmt"
	"learngo/github.com/chobkyu/hansik/config"
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
	for idx, loc := range locate {
		hansik := scrapper.Scrap(loc)
		fmt.Println(hansik)
		insertData(hansik, idx)
		//hansikArr = append(hansikArr, hansik)
	}

	fmt.Println(hansikArr)
	return c.JSON(http.StatusOK, nil)
}

func insertData(hansikData []scrapper.Hansikdang, locId int) *sql.Row {
	fmt.Println(len(hansikData))
	db := config.DB()

	for _, hansik := range hansikData {
		fmt.Println(hansik.Name)
		fmt.Println(hansik.Addr)
		fmt.Println(hansik.Star)
		fmt.Println("???")

		sqlStatement := `insert into public.hansic (name,addr,star,locationId)
						values ($1,$2,$3,$4) returning id`

		fmt.Println("!!!")

		err := db.Raw(sqlStatement, hansik.Name, hansik.Addr, hansik.Star, locId+1)

		if err != nil {
			fmt.Println(&err)
			sqlStatement := `update public.hansic name = $1, star $2 where addr = $3`
			db.Raw(sqlStatement, hansik.Name, hansik.Star, hansik.Addr)
		}

	}
	//db.Close()

	return nil

}
