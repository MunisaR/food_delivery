package routes

import (
	"food_delivery/middleware"
	"headfirstgo/food_delivery/controllers"

	cors "github.com/rs/cors/wrapper/gin"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.Use(cors.AllowAll())
	/////translation
	// r.GET("/:locale", controllers.Translation)

	/////auth routes/////////
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)


	//////products routes///////
	r.GET("/products", controllers.FindProducts)
	r.GET("/products/:id", controllers.FindProductById)
	r.GET("/productbycategory/:category_id", controllers.FindProductByCategoryId)
	r.POST("/product", controllers.AddProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	// r.Use(middleware.Authentication())
	//////////users routes///////////

	r.Use(middleware.Authentication())

	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/:users/:id", controllers.UpdateUser)
	r.DELETE("/:users/:id", controllers.DeleteUser)



	///category routes/////////
	r.POST("/createCategory", controllers.CreateCategory)
	r.GET("/getAllCategories", controllers.GetAllCategories)
	r.GET("/categories/:id", controllers.GetCategoryById)
	r.PATCH("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)

	////cart routes/////////
	// r.GET("/getBasket", controllers.GetBasket)
	// r.POST("/addBasket", controllers.AddNewBasket)
	r.PUT("/addToBasket", controllers.AddItemsToBasket)

	return r
}
