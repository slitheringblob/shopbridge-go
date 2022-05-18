package main

import (
	"shopbridge-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := models.SetUpDB()
	db.AutoMigrate(&models.Item{})
	router.Run(":8080")
}
