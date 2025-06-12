package routes

import (
	controllercategory "go_gin/controller/ControllerCategory"
	controllerproduct "go_gin/controller/ControllerProduct"
	controlleruser "go_gin/controller/ControllerUser"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	r.GET("/product/all", controllerproduct.GetAllProduct)
	r.GET("/product/all/filter", controllerproduct.GetProduct)
	r.POST("/product/insert", controllerproduct.InsertProduct)
	r.PUT("/product/update", controllerproduct.UpdateProduct)

	r.GET("/user/all", controlleruser.GetAllUser)
	r.GET("/user/all/filter", controlleruser.GetUser)
	r.POST("/user/insert", controlleruser.InsertUser)

	r.GET("/category/all", controllercategory.GetAllCategory)
	r.POST("/category/insert", controllercategory.InsertCategory)
}
