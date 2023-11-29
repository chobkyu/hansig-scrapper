package repositories

import (
	"learngo/github.com/chobkyu/hansik/models"
	"learngo/github.com/chobkyu/hansik/storage"
)

func InsertData(test models.Test) (models.Test, error) {
	db := storage.GetDB()

	sqlStatement := `insert into hansic (name, addr, star) 
					values ($1,$2,$3) RETIRNING id`
	err := db.QueryRow(sqlStatement, test.Name, test.Addr, test.Star).Scan(&test.Id)

	if err != nil {
		return test, err
	}

	return test, nil
}
