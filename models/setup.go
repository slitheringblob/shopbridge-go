package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB() *gorm.DB {
	user := "postgres"
	password := "postgres"
	host := "localhost"
	port := "5433"
	dbname := "shopbridge_go"

	db_url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	db, err := gorm.Open(postgres.Open(db_url))
	if err != nil {
		panic(err.Error())
	}
	return db
}
