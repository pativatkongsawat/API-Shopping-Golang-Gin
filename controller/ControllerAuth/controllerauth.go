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

func GenerateToken(user users.Users) (string, error) {
	claims := users.AuthClaims{
		UserId: user.ID,
		Email:  user.Email,
		Role:   user.PermissionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
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

		// ✅ ตรวจสอบก่อน hash
		if !helper.IsValidPassword(input.Password) {
			ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
				Status:  400,
				Message: "Invalid password format",
				Result:  "Password must include uppercase, lowercase, number, and special character",
			})
			return
		}

		if !helper.IsValidNameFormat(input.Firstname) {
			ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
				Status:  400,
				Message: "Invalid firstname format",
				Result:  "Firstname must not contain numbers and must start with uppercase",
			})
			return
		}

		if !helper.IsValidNameFormat(input.Lastname) {
			ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{
				Status:  400,
				Message: "Invalid lastname format",
				Result:  "Lastname must not contain numbers and must start with uppercase",
			})
			return
		}

		// ✅ Hash หลังจากผ่านการตรวจสอบ
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
			ID:           helper.GenerateUUID(),
			Firstname:    input.Firstname,
			Lastname:     input.Lastname,
			Address:      input.Address,
			Email:        input.Email,
			Password:     string(hashedPassword),
			CreatedAt:    &now,
			UpdatedAt:    &now,
			DeletedAt:    nil,
			PermissionID: 2,
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

	ctx.JSON(http.StatusCreated, utils.ResponseMessage{
		Status:  http.StatusCreated,
		Message: "REGISTER SUCCESS",
		Result:  user,
	})
}

func Login(ctx *gin.Context) {

	var input users.UserLogin

	if err := ctx.ShouldBindJSON(&input); err != nil {

		ctx.JSON(http.StatusBadRequest, utils.ResponseMessage{

			Status:  400,
			Message: "BAD REQUEST",
			Result:  err.Error(),
		})
		return

	}

	db := database.DBMYSQL

	var user users.Users

	if err := db.Where("email = ? ", input.Email).First(&user).Error; err != nil {

		ctx.JSON(http.StatusUnauthorized, utils.ResponseMessage{
			Status:  401,
			Message: "INVALID EMAIL OR PASSWORD",
			Result:  err.Error(),
		})
		return

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {

		ctx.JSON(http.StatusUnauthorized, utils.ResponseMessage{
			Status:  401,
			Message: "INVALID EMAIL OR PASSWORD",
			Result:  err.Error(),
		})
		return

	}
	token, err := GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "FAILED TO GENERATE TOKEN",
			Result:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "LOGIN SUCCESS",
		Result: gin.H{
			"token":        token,
			"user_id":      user.ID,
			"email":        user.Email,
			"firstname":    user.Firstname,
			"lastname":     user.Lastname,
			"permissionID": user.PermissionID,
		},
	})

}
