package v2handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (ph *ProductHandler) GetProductsV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of Products (V2)",
	})
}

func (ph *ProductHandler) GetproductByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID (V2)",
	})
}

func (ph *ProductHandler) PostProductsV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created (V2)",
	})
}

func (ph *ProductHandler) PutProductsByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product updated (V2)",
	})
}

func (ph *ProductHandler) DeleteProductsV2(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Product deleted (V2)",
	})
}
