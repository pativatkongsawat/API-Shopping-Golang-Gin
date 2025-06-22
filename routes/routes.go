package routes

import (
	controllerauth "go_gin/controller/ControllerAuth"
	controllercategory "go_gin/controller/ControllerCategory"
	controllerorder "go_gin/controller/ControllerOrder"
	controllerproduct "go_gin/controller/ControllerProduct"
	controlleruser "go_gin/controller/ControllerUser"
	"go_gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Auth
	api.POST("/auth/register", controllerauth.Register)
	api.POST("/auth/login", controllerauth.Login)

	// Users
	api.GET("/users", controlleruser.GetAllUser)
	api.GET("/users/search", controlleruser.GetUser)
	api.POST("/users", controlleruser.InsertUser)

	// Products
	api.GET("/products", controllerproduct.GetAllProduct)
	api.GET("/products/search", controllerproduct.GetProduct)
	api.POST("/products", controllerproduct.InsertProduct)
	api.PUT("/products", controllerproduct.UpdateProduct)
	api.DELETE("/products", controllerproduct.DeleteProduct)

	// Categories
	api.GET("/categories", middleware.AuthMiddleware(), controllercategory.GetAllCategory)
	api.POST("/categories", middleware.AuthMiddleware(), controllercategory.InsertCategory)
	api.PUT("/categories", middleware.AuthMiddleware(), controllercategory.UpdateCategory)
	api.DELETE("/categories", middleware.AuthMiddleware(), controllercategory.DeleteCategory)

	//Order
	api.POST("/order", middleware.AuthMiddleware(), controllerorder.CreateOrder)
}
