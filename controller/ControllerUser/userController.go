package controlleruser

import (
	"go_gin/database"
	"go_gin/helper"
	"go_gin/models/users"
	"go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all User
// @Description Get all User from the database
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} users.Users
// @Router /user/all [get]
func GetAllUser(ctx *gin.Context) {
	usermodelhelper := users.UserModelHelper{DB: database.DBMYSQL}

	users, err := usermodelhelper.GetAllUser()
	if err != nil {
		ctx.JSON(500, utils.ResponseMessage{
			Status:  http.StatusInternalServerError,
			Message: "Error getting users",
			Result:  err.Error(),
		})
		return
	}

	if len(users) == 0 {
		ctx.JSON(200, gin.H{
			"user":    users,
			"Message": "No users found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"user":    users,
		"Message": "Success",
	})
}

func GetUser(ctx *gin.Context) {
	usermodelhelper := users.UserModelHelper{DB: database.DBMYSQL}
	var userfil helper.UserFilter

	if err := ctx.ShouldBindQuery(&userfil); err != nil {
		ctx.JSON(400, utils.ResponseMessage{
			Status:  400,
			Message: "Error getting user",
			Result:  err.Error(),
		})
		return
	}

	user, count, err := usermodelhelper.GetUser(userfil.Fname, userfil.Lname, userfil.Email, userfil.Limit, userfil.Page)
	if err != nil {
		ctx.JSON(500, utils.ResponseMessage{
			Status:  500,
			Message: "Error getting user",
			Result:  err.Error(),
		})
		return
	}

	totalPages := (count + int64(userfil.Limit) - 1) / int64(userfil.Limit)
	nextPage := userfil.Page + 1
	if nextPage > int(totalPages) {
		nextPage = int(totalPages)
	}
	prevPage := userfil.Page - 1
	if prevPage < 1 {
		prevPage = 1
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
		"user":    user,
		"Pagination": helper.Pagination{
			Totalrows:     int(count),
			Totalpage:     int(totalPages),
			TotalNextpage: nextPage,
			Totalprevpage: prevPage,
			Prevpage:      prevPage,
			Nextpage:      nextPage,
		},
	})
}

func InsertUser(ctx *gin.Context) {

}
