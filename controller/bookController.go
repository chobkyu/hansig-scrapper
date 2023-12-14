package controller

import (
	"fmt"
	"learngo/github.com/chobkyu/hansik/config"
	"learngo/github.com/chobkyu/hansik/models"
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

func CreateBook(c echo.Context) error {
	fmt.Println("create")
	b := new(models.Book)
	db := config.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	book := &models.Book{
		Name:        b.Name,
		Description: b.Description,
	}

	if err := db.Create(&book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateBook(c echo.Context) error {
	id := c.Param(("id"))
	b := new(models.Book)
	db := config.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_book := new(models.Book)

	if err := db.First(&existing_book, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_book.Name = b.Name
	existing_book.Description = b.Description
	if err := db.Save(&existing_book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_book,
	}

	return c.JSON(http.StatusOK, response)
}

func GetBook(c echo.Context) error {
	fmt.Println("get")
	id := c.Param("id")
	db := config.DB()

	fmt.Println(id)

	var books []*models.Book

	if res := db.Find(&books, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": books,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	book := new(models.Book)

	err := db.Delete(&book, id).Error

	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a book has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}

func CreateData(c echo.Context) error {
	fmt.Println("get Data")
	h := new(models.Hansic)
	//db := config.DB()

	if err := c.Bind(h); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	getHansikData()

	return c.JSON(http.StatusOK, true)

}

func getHansikData() {
	var hansikArr []hansikdang
	for idx, loc := range locate {
		hansik := scrapper.Scrap(loc)
		fmt.Println(hansik)
		insertData(hansik, idx)
		//hansikArr = append(hansikArr, hansik)
	}
	fmt.Println(hansikArr)
}

func insertData(hansikData []scrapper.Hansikdang, idx int) {
	//h := new(models.Test)
	db := config.DB()

	for _, hansik := range hansikData {
		data := &models.Hansic{
			Name:       hansik.Name,
			Addr:       hansik.Addr,
			GoogleStar: hansik.Star,
			LocationId: idx + 1,
		}

		if err := db.Create(&data).Error; err != nil {
			data := map[string]interface{}{
				"message": err.Error(),
			}

			fmt.Println(data)
		}
	}
}
