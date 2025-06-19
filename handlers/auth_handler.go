package handlers

import (
	"final-golang-project/services"
	utils "final-golang-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
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

	err := h.service.RegisterUser(request.Username, request.Email, request.Password)
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

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "validation error",
		})
		return
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := utils.GenerateJwt(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *AuthHandler) GetUserByEmail(ctx *gin.Context) {
	email := ctx.GetString("email")
	user, error := h.service.GetUserByEmail(email)
	if error != nil {
		ctx.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	if user == nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(200, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
	})
}
