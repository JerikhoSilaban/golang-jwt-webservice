package router

import (
	"DTSGolang/Kelas3/Sesi2Bagian2/controllers"
	"DTSGolang/Kelas3/Sesi2Bagian2/middlewares"

	"github.com/gin-gonic/gin"

	_ "DTSGolang/Kelas3/Sesi2Bagian2/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title JWT Product CRUD API with Authentication and Authorization
// @version 1.0
// @description This is a simple service for managing products by admin and users
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8000
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())

		// Post product
		productRouter.POST("/", controllers.CreateProduct)

		// Update product
		productRouter.PUT("/:productID", middlewares.ProductAuthorization(), controllers.UpdateProduct)

		// Get product by id
		productRouter.GET("/:productID", middlewares.ProductAuthorization(), controllers.GetProductById)

		// Get all products
		productRouter.GET("/", middlewares.ProductAuthorizationAll(), controllers.GetProducts)

		// Delete product by id
		productRouter.DELETE("/:productID", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
