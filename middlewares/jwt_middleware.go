package middlewares

import (
	utils "final-golang-project/utlis"
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
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is invalid"})
			ctx.Abort()
			return
		}

		// authHeader = "Bearer abchdhashd328923hdajsdn.sdfsdfh.sdhfsfh"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, error := utils.ValidateJwt(tokenString)
		if error != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println("Claims:", claims)
		ctx.Set("email", claims["email"])

		ctx.Next()
	}
}
