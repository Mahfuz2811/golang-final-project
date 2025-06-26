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

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
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

	fmt.Println("Starting server on :8080...")
	if error := router.Run(":8080"); error != nil {
		fmt.Println("Error of starting server: ", error)
	}
}
