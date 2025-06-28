package main

import (
	"final-golang-project/db"
	"final-golang-project/handlers"
	"final-golang-project/middlewares"
	"final-golang-project/rabbitmq"
	"final-golang-project/redis"
	"final-golang-project/repositories"
	"final-golang-project/routes"
	"final-golang-project/services"
	"final-golang-project/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode from environment
	ginMode := getEnv("GIN_MODE", "release")
	gin.SetMode(ginMode)

	database, error := db.NewMySqlDB()
	if error != nil {
		panic(error)
	}

	gormDb, error := db.NewMySqlGormDB()
	if error != nil {
		panic(error)
	}

	redisClient, error := redis.NewRedisClient()
	if error != nil {
		panic(error)
	}

	rabbitmqClient, error := rabbitmq.NewRabbitMQ()
	if error != nil {
		panic(error)
	}

	emailSender := utils.NewEmailSender(rabbitmqClient)

	userRepo := repositories.NewMySQLUserRepository(database)
	redisMySQLUserRepo := repositories.NewRedisMySQLUserRepository(userRepo, redisClient)
	authService := services.NewAuthServe(redisMySQLUserRepo, emailSender)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	// Health check endpoint for Docker
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	router.GET("/user", middlewares.JWTAuthMiddleware(), authHandler.GetUserByEmail)

	// migrate product table
	// if error := db.MigrateProductTable(gormDb); error != nil {
	// 	panic(error)
	// }
	//product route handler & dependency
	productRepo := repositories.NewProductRepository(gormDb)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	routes.RegisterProductRoutes(router, productHandler)

	port := getEnv("APP_PORT", "8080")
	fmt.Printf("Starting server on :%s...\n", port)
	if error := router.Run(":" + port); error != nil {
		fmt.Println("Error of starting server: ", error)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
