package controllercategory

import (
	"go_gin/database"
	"go_gin/models/category"
	"go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertCategory(ctx *gin.Context) {

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

	categorydata := []category.Category{}

	if err := ctx.ShouldBindJSON(&categorydata); err != nil {

		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{

			Status:  400,
			Message: "CATEGORY BAD REQUEST",
			Result:  err.Error(),
		})

	}

	category, err := categoryModelHelper.InsertCategory(categorydata)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "ERROR INSERT CATEGORY",
			Result:  err.Error(),
		})

	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "INSERT CATEGORYSUCSSFULY",
		Result:  category,
	})

}

func GetAllCategory(ctx *gin.Context) {

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

	categorys, err := categoryModelHelper.GetAllCategory()

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "ERROR FOR GET ALL CATEGORY",
			Result:  err.Error(),
		})

	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "SUCCESFULY",
		Result:  categorys,
	})

}
