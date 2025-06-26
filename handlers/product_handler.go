package handlers

import (
	"final-golang-project/models"
	"final-golang-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (p *ProductHandler) Create(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	userEmail := ctx.GetString("email")
	if userEmail == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})

		return
	}

	product.UserEmail = userEmail

	if err := p.service.Create(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erroe": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully.",
	})
}

func (p *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// product

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
