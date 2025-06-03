package main

import (
	"final-golang-project/db"
	"final-golang-project/handlers"
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

	userRepo := repositories.NewUserRepositoy(database)
	authService := services.NewAuthServe(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/register", authHandler.Register)

	fmt.Println("Starting server on :8080...")
	if error := router.Run(":8080"); error != nil {
		fmt.Println("Error of starting server: ", error)
	}
}
