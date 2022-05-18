package routes

import (
	"shopbridge-go/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetUpRoutes(db *gorm.DB) *gin.Engine {

	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	router.GET("/items", controllers.GetAllItems)
	router.POST("/items", controllers.CreateSingleItem)
	router.GET("/items/:id", controllers.GetSingleItem)
	router.PATCH("/items:id", controllers.UpdateSingleItem)
	router.DELETE("/items/:id", controllers.DeleteSingleItem)

	return router
}
