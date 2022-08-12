package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddProductInput struct {
	gorm.Model
	Title        string `gorm:"column:title" json:"title"`
	Description  string `gorm:"column:description" json:"description"`
	Price        uint    `gorm:"column:price"`
	Image        string `gorm:"column:image" json:"image"`
	CategoryID   int    `gorm:"column:category_id;foreignkey:product_id" json:"category_id"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}
type UpdateProductInput struct {
	gorm.Model
	Title        string `gorm:"column:title" json:"title"`
	Description  string `gorm:"column:description" json:"description"`
	Price        uint    `gorm:"column:price" json:"price"`
	Image        string `gorm:"column:image" json:"image"`
	CategoryID   int    `gorm:"column:category_id;foreignkey:product_id" json:"category_id"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

//////Find All Products
func FindProducts(c *gin.Context) {
	// var pagination controllers.GeneratePagination(c)
	// offset := (pagination.Page - 1) * pagination.Limit   .Limit(pagination.Limit).Offset(offset)

	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	if err := db.Scopes(Paginate(c)).Order("category_id asc").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Something went wrong",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

//////Find Products By its Id/////
func FindProductById(c *gin.Context) {
	//get model if exists
	var product []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      "Record not found",
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

///// Find Products By Category Id
func FindProductByCategoryId(c *gin.Context) {
	var products []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Scopes(Paginate(c)).Where("category_id = ?", c.Param("category_id")).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:category_id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

//////Adding Product
func AddProduct(c *gin.Context) {
	//validate input
	var input AddProductInput
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/product not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	

	//Create product
	product := models.Product{Title: input.Title, Description: input.Description, Price: input.Price, Image: input.Image, CategoryID: input.CategoryID, CategoryName: input.CategoryName}
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

///Updating Product
func UpdateProduct(c *gin.Context) {
	var input 	UpdateProductInput
	//Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	/////Updating ProductInputs
	var updateInput models.Product
	updateInput.Title = input.Title
	updateInput.Description = input.Description
	updateInput.Price = input.Price
	updateInput.Image = input.Image
	updateInput.CategoryID = input.CategoryID
	updateInput.CategoryName = input.CategoryName

	db.Model(&product).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func GetProductsExceptPizza(c *gin.Context){
	var products []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Scopes(Paginate(c)).Order("category_id ASC").Where("category_id != ?", 1).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:category_id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
	
}

/////Deleting Products
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	///get model if exists
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})

}
