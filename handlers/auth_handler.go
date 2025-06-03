package handlers

import (
	"final-golang-project/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Pasword  string `json:"password" binding:"required,min=6"`
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(422, gin.H{"error": "Validation error"})
		return
	}

	err := h.service.RegisterUser(request.Username, request.Email, request.Pasword)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to register",
		})

		return
	}

	ctx.JSON(201, gin.H{
		"message": "User registered successfully",
	})
}

// func (h *AuthHandler) GetUserByEmail(ctx *gin.Context) {

// }
