package main

import (
	"headfirstgo/food_delivery/database"
	"headfirstgo/food_delivery/models"
	"headfirstgo/food_delivery/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//close databse when the main func is finishes
	db := database.SetupPostgres()
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	r := routes.UserRoutes(db)
	routes.ProductRoutes(db)
	r.Run()
}
