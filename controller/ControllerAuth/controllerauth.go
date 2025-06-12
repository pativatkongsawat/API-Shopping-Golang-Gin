package controllerauth

import (
	"go_gin/database"
	"go_gin/helper"
	"go_gin/models/users"
	"net/http"
	"os"
	"time"

	"go_gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(email string) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	expTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(expTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Register(ctx *gin.Context) {
	now := time.Now()

	var userInputs []users.UsersInsert

	if err := ctx.ShouldBindJSON(&userInputs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
			Status:  400,
			Message: "BAD REQUEST",
			Result:  err.Error(),
		})
		return
	}

	var usersToCreate []users.Users
	for _, input := range userInputs {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
				Status:  500,
				Message: "FAILED TO HASH PASSWORD",
				Result:  err.Error(),
			})
			return
		}

		user := users.Users{
			ID:        helper.GenerateUUID(),
			Firstname: input.Firstname,
			Lastname:  input.Lastname,
			Address:   input.Address,
			Email:     input.Email,
			Password:  string(hashedPassword),
			CreatedAt: &now,
			UpdatedAt: &now,
			DeletedAt: nil,
		}
		usersToCreate = append(usersToCreate, user)
	}

	userModelHelper := users.UserModelHelper{DB: database.DBMYSQL}

	user, err := userModelHelper.Register(usersToCreate)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "FAIL TO REGISTER",
			Result:  err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "REGISTER SUCCESS",
		Result:  user,
	})
}

func Login(ctx *gin.Context) {

}
