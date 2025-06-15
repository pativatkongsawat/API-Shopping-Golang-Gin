package controllercategory

import (
	"go_gin/database"
	"go_gin/models/category"
	"go_gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertCategory(ctx *gin.Context) {

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

	categorydata := []category.Category{}

	if err := ctx.ShouldBindJSON(&categorydata); err != nil {

		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{

			Status:  400,
			Message: "ERROR BIND DATA",
			Result:  err.Error(),
		})
		return

	}

	category, err := categoryModelHelper.InsertCategory(categorydata)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "ERROR INSERT CATEGORY",
			Result:  err.Error(),
		})
		return

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
		return

	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "SUCCESFULY",
		Result:  categorys,
	})

}

func DeleteCategory(ctx *gin.Context) {

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

	idstr := ctx.Query("id")

	if idstr == "" {

		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{

			Status:  400,
			Message: "BAD REQUEST",
			Result:  idstr,
		})
		return

	}

	id, err := strconv.Atoi("idstr")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
			Status:  400,
			Message: "BAD REQUEST",
			Result:  err.Error(),
		})
		return
	}

	category, err := categoryModelHelper.DeleteCategory(id)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "CAN'T DELETE CATEGORY",
			Result:  err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "DELETE SUCCESSFULY",
		Result:  category,
	})

}

func UpdateCategory(ctx *gin.Context) {

	categoryModelHelper := category.CategoryModelHelper{DB: database.DBMYSQL}

	var categorydata []category.Category

	if err := ctx.ShouldBindJSON(&categorydata); err != nil {

		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
			Status:  400,
			Message: "ERROR BIND DATA",
			Result:  err.Error(),
		})
		return
	}

	category, err := categoryModelHelper.UpdateCategory(categorydata)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "ERROR UPDATE CATEGORY",
			Result:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "UPDATE CATEGORY SUCESSFULY",
		Result:  category,
	})

}
