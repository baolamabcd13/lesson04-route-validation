package v1handler

import (
	"github.com/gin-gonic/gin"
	"lesson03-route-group/utils"
	"net/http"
)

type CategoryHandler struct{}

type PostCategoriesV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

type GetCategoryBycategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (ch *CategoryHandler) GetCategoryBycategoryV1(ctx *gin.Context) {
	var params GetCategoryBycategoryV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category details (V1)",
		"course":  params.Category,
	})
}
func (ch *CategoryHandler) PostCategoriesV1(ctx *gin.Context) {
	var params PostCategoriesV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post Category (V1)",
		"name":    params.Name,
		"status":  params.Status,
	})
}
