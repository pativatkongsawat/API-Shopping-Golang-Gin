package controllerproduct

import (
	"go_gin/database"
	"go_gin/helper"
	"go_gin/models/products"
	"go_gin/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get all products
// @Description Get all products from the database
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {array} products.Product
// @Router /product/all [get]
func GetAllProduct(ctx *gin.Context) {
	productmodelhelper := products.ProductModelHelper{DB: database.DBMYSQL}
	products, err := productmodelhelper.GetAllProducts()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "Success",
		"products": products})
}

func GetProduct(ctx *gin.Context) {

	productmodelhelper := products.ProductModelHelper{DB: database.DBMYSQL}

	var lipage helper.LimitPage

	if err := ctx.ShouldBindQuery(&lipage); err != nil {
		ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Failed to bind query",
			Result:  err.Error(),
		})
	}
	if lipage.Limit <= 0 {
		lipage.Limit = 10
	}

	if lipage.Page <= 0 {
		lipage.Page = 1
	}

	product, count, err := productmodelhelper.GetProduct(lipage.PName, lipage.Limit, lipage.Page)

	if err != nil {
		ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Failed to get product",
			Result:  err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"product": product,
		"Pagination": helper.Pagination{
			Totalrows:     int(count),
			Totalpage:     int(count) / lipage.Limit,
			Prevpage:      lipage.Page - 1,
			Nextpage:      lipage.Page + 1,
			TotalNextpage: (int(count) / lipage.Limit) - lipage.Page,
			Totalprevpage: lipage.Page - 1,
		},
	})
}
