package main

import (
	"final-golang-project/db"
	"final-golang-project/handlers"
	"final-golang-project/redis"
	"final-golang-project/repositories"
	"final-golang-project/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	database, error := db.NewMySqlDB()
	if error != nil {
		panic(error)
	}

	redisClient, error := redis.NewsRedisClient()
	if error != nil {
		panic(error)
	}

	userRepo := repositories.NewMySQLUserRepositoy(database)
	redisMySQLUserRepo := repositories.NewRedisMySQLUserRepository(userRepo, redisClient)
	authService := services.NewAuthServe(redisMySQLUserRepo)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/register", authHandler.Register)
	router.GET("/user", authHandler.GetUserByEmail)

	fmt.Println("Starting server on :8080...")
	if error := router.Run(":8080"); error != nil {
		fmt.Println("Error of starting server: ", error)
	}
}
