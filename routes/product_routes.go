package routes

import (
	"final-golang-project/handlers"
	"final-golang-project/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, handler *handlers.ProductHandler) {
	authGroup := router.Group("/products")
	authGroup.Use(middlewares.JWTAuthMiddleware())

	authGroup.POST("/create", handler.Create)
	authGroup.GET("/list", handler.List)
	// authGroup.Use(middlewares.JWTAuthMiddleware())
	// {
	// 	authGroup.POST("/create", handler.Create)
	// }
}
