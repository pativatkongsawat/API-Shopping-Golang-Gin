package routes

import (
	controllerauth "go_gin/controller/ControllerAuth"
	controllercategory "go_gin/controller/ControllerCategory"
	controllerproduct "go_gin/controller/ControllerProduct"
	controlleruser "go_gin/controller/ControllerUser"

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
	api.GET("/categories", controllercategory.GetAllCategory)
	api.POST("/categories", controllercategory.InsertCategory)
	api.PUT("/categories", controllercategory.UpdateCategory)
	api.DELETE("/categories", controllercategory.DeleteCategory)
}
