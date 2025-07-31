package v1handler

import (
	"github.com/gin-gonic/gin"
	"lesson03-route-group/utils"
	"net/http"
	"time"
)

type ProductHandler struct{}

type GetProductBySLugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=5,max=100"`
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=5,max=100,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"` //year-month-day
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (ph *ProductHandler) GetProductsV1(ctx *gin.Context) {
	var params GetProductsV1Param
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))

		return
	}

	if params.Limit == 0 {
		params.Limit = 1
	}

	if params.Email == "" {
		params.Email = "No email"
	}

	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of Products (V1)",
		"search":  params.Search,
		"limit":   params.Limit,
		"email":   params.Email,
		"date":    params.Date,
	})
}

func (ph *ProductHandler) GetProductBySLugV1(ctx *gin.Context) {
	var params GetProductBySLugV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID (V1)",
		"slug":    params.Slug,
	})
}

func (ph *ProductHandler) PostProductsV1(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error read body request")
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created (V1)",
		"data":    string(body),
	})
}

func (ph *ProductHandler) PutProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product updated (V1)",
	})
}

func (ph *ProductHandler) DeleteProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Product deleted (V1)",
	})
}
