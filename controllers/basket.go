package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddBasketInput struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

type UpdateBasketInput struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

type UpdateBasketItemInput struct {
	gorm.Model
	// BasketID  uint `json:"basket_id"`
	Items []models.BasketItem
	// ProductID uint `json:"product_id"`
}

func GetBasket(c *gin.Context) {
	var basket models.Basket
	var user_id = c.GetInt("id")
	////checking if user_id exists
	// if !user_id_exists {
	// 	c.JSON(401, gin.H{"message": "user_id not found"})
	// }
	fmt.Println(user_id)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", user_id).Preload("Item").Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: user_id, TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 200,
			"data":       newBasket,
		})
		return
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message":    "Route GET:/getAllCategories not found",
		// 	"error":      "Record not found",
		// 	"statusCode": 404,
		// })
		// return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})
}

func AddItem(c *gin.Context) {
	var basket_item models.BasketItem
	var input AddBasketInput
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route PUT:/basket not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	//Create product
	basketItem := models.BasketItem{BasketID: basket_item.BasketID, ProductID: basket_item.ProductID, Quantity: basket_item.Quantity}
	db.Create(&basketItem)
	c.JSON(http.StatusOK, gin.H{"data": basketItem})

}

func UpdateBasket(c *gin.Context) {

	// var products models.Product
	var basket models.Basket
	var basketItems []models.BasketItem
	// var user models.User
	var total_price uint
	var updateBasketItemInput UpdateBasketItemInput

	var user_id = c.GetInt("id")
	paramInt, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	// user_idInt := int(user_id)
	////checking if user_id exists
	// var basket = find({user_id: c.id})
	db := c.MustGet("db").(*gorm.DB)
	//// find basket by user_id_exists
	if err := db.Where("id = ?", paramInt).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})

		return
	}
	if basket.UserID != user_id {
		c.JSON(http.StatusForbidden, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": http.StatusForbidden,
		})
		return
	}

	fmt.Println(paramInt)
	if err := db.Where("basket_id = ?", basket.ID).Find(&basketItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/basketitems not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	fmt.Println(basketItems)
	var aa map[uint]models.BasketItem

	for _, item := range basketItems {
		aa[item.ID] = item
	} 

	var res []models.BasketItem
	for _, item := range updateBasketItemInput.Items {
		if val, ok := aa[item.ID]; ok {
			aa[item.ID] = val
			total_price += item.Quantity * item.Product.Price
		} else {
			aa[item.ID] = item
			total_price += item.Quantity * item.Product.Price
		}
		res = append(res, aa[item.ID])
	}

	db.Model(&basketItems).Updates(res)
	fmt.Println(user_id)
	fmt.Println(c.Get("phone_number"))
	fmt.Println(c.Get("id"))

	c.JSON(http.StatusOK, gin.H{"data": basketItems})
}

func GetBasketById(c *gin.Context) {
	var basket []models.Basket
	var user_id = c.GetInt("id")
	// user_idInt := int(user_id)
	////checking if user_id exists

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: user_id, TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"message":    "Created a new basket",
			"error":      "",
			"statusCode": 200,
			"data":       newBasket,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
