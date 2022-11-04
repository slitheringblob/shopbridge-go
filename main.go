package main

import (
	"log"
	"shopbridge-go/models"
	"shopbridge-go/routes"
)

func main() {

	db := models.SetUpDB()
	db.AutoMigrate(&models.Item{})

	router := routes.SetUpRoutes(db)
	log.Fatal(router.Run(":8080"))
}
