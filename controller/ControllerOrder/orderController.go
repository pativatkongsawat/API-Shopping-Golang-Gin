package controllerorder

import (
	"go_gin/database"
	"go_gin/models/order"
	"go_gin/models/users"
	"go_gin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	ordermodelhelper := order.OrderModelHelper{DB: database.DBMYSQL}

	data := order.OrderCreateRequest{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
			Status:  400,
			Message: "Error Bind Data",
			Result:  err.Error(),
		})
		return
	}

	claimAny, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ResponseMessage{
			Status:  401,
			Message: "Unauthorized",
			Result:  nil,
		})
		return
	}

	claim := claimAny.(*users.AuthClaims)
	now := time.Now()

	var totalPrice float64 = 0
	for _, item := range data.Products {
		totalPrice += float64(item.Quantity) * item.Price
	}

	newOrder := order.Order{
		UserId:     claim.UserId,
		CreateAt:   &now,
		UpdatedAt:  &now,
		DeletedAt:  nil,
		TotalPrice: totalPrice,
		CreatedBy:  claim.Email,
	}

	order, err := ordermodelhelper.CreateOrder(&newOrder)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "Error Create order",
			Result:  err.Error(),
		})
		return
	}

	orderhas, err := ordermodelhelper.CreateOrderHasProduct(order.Id, data.Products)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "Error Create OrderHasProducts",
			Result:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "Create Order SuccessFuly",
		Result:  orderhas,
	})

}
