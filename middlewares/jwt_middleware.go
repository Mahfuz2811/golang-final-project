package middlewares

import (
	"final-golang-project/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization'header is missing",
			})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is invalid",
			})
			ctx.Abort()
			return
		}

		// Bearer eyJhbGciOsdhfjsdhfiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNzUwNDM1MzQ4fQ.et0fVXakIHCy8sdhfsdjfyGztyQqWX5NXPsh82VEGuVzXFSA6Sc
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, error := utils.ValidateJwt(tokenString)
		if error != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println("Claims:", claims)
		ctx.Set("email", claims["email"])

		ctx.Next()
	}
}
