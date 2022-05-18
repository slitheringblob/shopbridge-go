package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SetUpDB() *gorm.DB {
	user := "postgres"
	password := "postgres"
	host := "localhost"
	port := "5433"
	dbname := "shopbridge_go"

	db_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := gorm.Open("postgres", db_url)
	if err != nil {
		panic(err.Error())
	}
	return db
}
