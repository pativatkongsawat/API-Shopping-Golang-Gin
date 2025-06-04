package controllerproduct

import (
	"go_gin/database"
	"go_gin/helper"
	"go_gin/models/products"
	"go_gin/utils"
	"time"

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

// @Summary Get products
// @Description Get a list of products with pagination
// @Tags product
// @Accept  json
// @Produce  json
// @Param pname query string false "Product name"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {array} products.Product
// @Router /product/all/filter [get]
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

func InsertProduct(ctx *gin.Context) error {

	productmodelhelper := products.ProductModelHelper{DB: database.DBMYSQL}

	data := []products.InsertProduct{}

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(400, gin.H{
			"MESSAGE": "BAD REQUEST",
			"ERORR":   err.Error(),
		})
	}

	return nil
}

func UpdateProduct(ctx *gin.Context) {

	now := time.Now()

	productmodelhelper := products.ProductModelHelper{DB: database.DBMYSQL}
	productdata := []products.UpdateProduct{}

	newproduct := []products.Product{}

	if err := ctx.ShouldBindBodyWithJSON(&productdata); err != nil {
		ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Could not bind",
			Result:  err.Error(),
		})
	}

	for _, i := range productdata {
		newdata := products.Product{
			Id:          i.Id,
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
			Quantity:    i.Quantity,
			Update_at:   &now,
			Category_id: i.Category_id,
		}

		newproduct = append(newproduct, newdata)
	}

	product, err := productmodelhelper.UpdateProduct(newproduct)

	if err != nil {
		ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error updating product",
			Result:  err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message":  "Success",
		"Prouduct": product,
	})

}
