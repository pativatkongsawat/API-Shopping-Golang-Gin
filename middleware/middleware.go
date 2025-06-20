package middleware

import (
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
		claims := &users.AuthClaims{}

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

		ctx.Set("user", claims)
		ctx.Set("user_email", claims.Email)
		ctx.Set("user_role", claims.Role)
		ctx.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		role, _ := ctx.Get("user_role")

		if role != "Admin" {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{

				"Error": "Admin only",
			})
			return

		}
		ctx.Next()

	}
}

func CustomerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		role, _ := ctx.Get("user")

		if role != "" {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{

				"Error": "Customer only",
			})
			return

		}
		ctx.Next()

	}
}
