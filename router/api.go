package router

import (
	"database/sql"
	"fmt"
	"learngo/github.com/chobkyu/hansik/scrapper"
	"learngo/github.com/chobkyu/hansik/storage"
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
		insertData(hansik)
		//hansikArr = append(hansikArr, hansik)
	}

	fmt.Println(hansikArr)
	return c.JSON(http.StatusOK, nil)
}

func insertData(hansikData []scrapper.Hansikdang) *sql.Row {
	fmt.Println(len(hansikData))
	db := storage.GetDB()

	for _, hansik := range hansikData {
		fmt.Println(hansik.Name)
		fmt.Println(hansik.Addr)
		fmt.Println(hansik.Star)

		sqlStatement := `insert into test.hansic (name,addr,star)
						values ($1,$2,$3) returning id`

		err := db.QueryRow(sqlStatement, hansik.Name, hansik.Addr, hansik.Star)

		if err != nil {
			sqlStatement := `update test.hansic name = $1, star $2 where addr = $3`
			db.QueryRow(sqlStatement, hansik.Name, hansik.Star, hansik.Addr)
		}

	}
	db.Close()

	return nil

}
