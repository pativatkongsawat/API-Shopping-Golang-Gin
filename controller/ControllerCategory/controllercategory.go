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
