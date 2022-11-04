package controllers

import (
	"fmt"
	"net/http"
	"shopbridge-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//################################
//Schemas to Validate USers Input
//################################
type itemUserSchema struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price,string"`
}

//controller functions for routes

//GET /item
func GetAllItems(c *gin.Context) {
	fmt.Println("GetAllItems Controller")
	db := c.MustGet("db").(*gorm.DB)
	var allItems []models.Item
	db.Find(&allItems)

	c.JSON(http.StatusOK, gin.H{"data": allItems})

}

//POST /item
func CreateSingleItem(c *gin.Context) {

	fmt.Println("CreateSingleItem Controller")

	var input models.Item

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db := c.MustGet("db").(*gorm.DB)
	name := c.Query("name")
	description := c.Query("description")
	price, err := strconv.ParseFloat(c.Query("price"), 64)

	if err != nil {
		fmt.Println("Error creating Item, ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {

		input.Name = name
		input.Description = description
		input.Price = price
		db.Create(&input)
		c.JSON(http.StatusOK, gin.H{"data": "Added new Item " + (name)})
	}

}

//GET /item/:id
func GetSingleItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var item models.Item
	item_id, conv_err := strconv.ParseFloat(c.Param("id"), 64)
	if conv_err != nil {
		fmt.Println("Error While Converting Param to Integer")
	}
	result := db.First(&item, "id = ?", item_id)
	fmt.Println(result)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"data": "Error Retrieving Item. " + result.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": item})
	}

}

//PATCH /item/:id
func UpdateSingleItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var item models.Item
	id := c.Query("id")
	if err := db.Where("id = ?", id).First(&item).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	var input itemUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Error" + err.Error()})
		return
	}

	db.Model(&item).Updates(input)
	c.JSON(200, input)
}

// DELETE /item/:id
func DeleteSingleItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var item models.Item
	item_id, conv_err := strconv.ParseFloat(c.Param("id"), 64)

	if conv_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse the ID parameter, check the type fella!"})
	}
	if err := db.Where("id = ?", item_id).First(&item).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Delete(&item, item_id)
		c.JSON(http.StatusOK, gin.H{"Deleted": "Deleted item with ID: " + c.Param("id")})
	}
}
