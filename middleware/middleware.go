package middleware

import (
	"go_gin/database"
	"go_gin/models/users"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHender := ctx.GetHeader("Authorization")

		if authHender == "" || !strings.HasPrefix(authHender, "Bearer") {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"ERROR": "Unauthorized",
			})
			return

		}
		tokenStr := strings.TrimPrefix(authHender, "Bearer ")
		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}

		ctx.Set("user_email", claims.Subject)

		var user users.Users

		db := database.DBMYSQL

		if err := db.Where("email = ? ", claims.Subject).First(&user).Error; err != nil {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User Not Found ",
			})
			return

		}
		ctx.Set("user_role", user.PermissionID)

		ctx.Next()

	}

}
