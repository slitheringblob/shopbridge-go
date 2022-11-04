package routes

import (
	"shopbridge-go/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *gin.Engine {

	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	router.GET("/item", controllers.GetAllItems)       //done
	router.POST("/item", controllers.CreateSingleItem) //done
	router.GET("/item/:id", controllers.GetSingleItem) //done
	router.PATCH("/item/:id", controllers.UpdateSingleItem)
	router.DELETE("/item/:id", controllers.DeleteSingleItem)

	return router
}
