package v1handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lesson04-route-validation/utils"
	"net/http"
	"time"
)

type ProductHandler struct{}

type GetProductBySLugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=5,max=100"`
}

type ProductsImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required,file_ext=jpg png gif"`
}

type ProductsInfo struct {
	InfoKey   string `json:"info_key" binding:"required"`
	InfoValue string `json:"info_value" binding:"required"`
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=5,max=100,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"` //year-month-day
}

type ProductAttribute struct {
	AttributeName  string `json:"attribute_name" binding:"required"`
	AttributeValue string `json:"attribute_value" binding:"required"`
}

type PostProductsV1Param struct {
	Name             string                  `json:"name" binding:"required,min=3,max=100"`
	Price            int                     `json:"price" binding:"required,max_int=100000"`
	Display          *bool                   `json:"display" binding:"omitempty"`
	ProductImage     ProductsImage           `json:"product_image" binding:"required"`
	Tags             []string                `json:"tags" binding:"required,gt=3,lt=5"`
	ProductAttribute []ProductAttribute      `json:"product_attribute" binding:"required,gt=0,dive"`
	ProductsInfo     map[string]ProductsInfo `json:"products_info" binding:"required,gt=0,dive"`
	ProductMetadata  map[string]any          `json:"product_metadata" binding:"omitempty"`
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
	var params PostProductsV1Param
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))

		return
	}

	for key := range params.ProductsInfo {
		if _, err := uuid.Parse(key); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{
					"product_info": fmt.Sprintf("Key %s trong product_info không phải là UUID hợp lệ", key),
				},
			})
			return
		}
	}

	if params.Display == nil {
		defaultDisplay := true
		params.Display = &defaultDisplay
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":           "product created (V1)",
		"name":              params.Name,
		"price":             params.Price,
		"display":           params.Display,
		"product_image":     params.ProductImage,
		"tags":              params.Tags,
		"product_attribute": params.ProductAttribute,
		"product_info":      params.ProductsInfo,
		"product_metadata":  params.ProductMetadata,
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
