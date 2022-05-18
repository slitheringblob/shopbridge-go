package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	item_value := c.Param("id")
	name := c.Param("name")
	description := c.Param("description")
	price := c.Param("price")

}
func CreateSingleItem(c *gin.Context) {

}
func GetSingleItem(c *gin.Context) {

}
func UpdateSingleItem(c *gin.Context) {

}
func DeleteSingleItem(c *gin.Context) {

}
