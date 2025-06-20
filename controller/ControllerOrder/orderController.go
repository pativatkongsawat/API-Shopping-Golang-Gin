package controllerorder

import (
	"go_gin/models/order"
	"go_gin/models/users"
	"go_gin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {

	data := order.OrderCreateRequest{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
			Status:  400,
			Message: "Error Bind Data",
			Result:  err.Error(),
		})
		return
	}
	claimAny, err := ctx.Get("user")

	if !err {

		ctx.JSON(http.StatusUnauthorized, utils.ResponseMessage{
			Status:  401,
			Message: "Unauthorized",
			Result:  nil,
		})

	}

	claim := claimAny.(*users.AuthClaims)

	now := time.Now()

	newOrder := order.Order{
		UserId:     claim.UserId,
		CreateAt:   &now,
		UpdatedAt:  &now,
		DeletedAt:  nil,
		TotalPrice: 100,
		CreatedBy:  claim.Email,
	}

}
