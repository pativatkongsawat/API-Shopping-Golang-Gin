package middleware

import (
	"go_gin/database"
	"go_gin/models/users"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &jwt.RegisteredClaims{}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server misconfiguration"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid || claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		ctx.Set("user_email", claims.Subject)

		var user users.Users
		db := database.DBMYSQL

		if err := db.Where("email = ?", claims.Subject).First(&user).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User Not Found",
			})
			return
		}

		ctx.Set("user_role", user.PermissionID)
		ctx.Next()
	}
}
