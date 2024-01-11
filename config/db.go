package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB
var e error

func DatabaseInint() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_PORT"))
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASSWORD"))
	fmt.Println(os.Getenv("DB_NAME"))

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", dbHost, dbUser, dbPass, dbName, dbPort)

	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
