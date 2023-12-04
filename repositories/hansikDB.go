package repositories

import (
	"database/sql"
	"fmt"
	"learngo/github.com/chobkyu/hansik/models"
	"learngo/github.com/chobkyu/hansik/storage"
)

func InsertData(test models.Hansic) (models.Hansic, error) {
	db := storage.GetDB()

	fmt.Println(test.Name)
	fmt.Println(test.Addr)
	fmt.Println(test.Star)

	sqlStatement := `insert into test.hansic (name, addr, star) 
					values ($1,$2,$3) RETURNING id`
	err := db.QueryRow(sqlStatement, test.Name, test.Addr, test.Star).Scan(&test.Id)

	if err != nil {
		return test, err
	}

	return test, nil
}

func Test() (*sql.Rows, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM PG_TABLES`

	test, err := db.Query(sqlStatement)

	fmt.Println(&test)
	if err != nil {
		return test, err
	}
	return test, nil
}
