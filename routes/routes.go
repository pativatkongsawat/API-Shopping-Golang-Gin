package routes

import (
	controllerauth "go_gin/controller/ControllerAuth"
	controllercategory "go_gin/controller/ControllerCategory"
	controllerproduct "go_gin/controller/ControllerProduct"
	controlleruser "go_gin/controller/ControllerUser"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	//Product
	r.GET("/product/all", controllerproduct.GetAllProduct)
	r.GET("/product/all/filter", controllerproduct.GetProduct)
	r.POST("/product/insert", controllerproduct.InsertProduct)
	r.PUT("/product/update", controllerproduct.UpdateProduct)

	//User
	r.GET("/user/all", controlleruser.GetAllUser)
	r.GET("/user/all/filter", controlleruser.GetUser)
	r.POST("/user/insert", controlleruser.InsertUser)

	//Category
	r.GET("/category/all", controllercategory.GetAllCategory)
	r.POST("/category/insert", controllercategory.InsertCategory)

	//Auth
	r.POST("/register/", controllerauth.Register)
	r.POST("/login/" , controllerauth.Login)
}
