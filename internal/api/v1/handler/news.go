package v1handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsHandler struct {
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) GetNewsV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get news (V1)",
			"slug":    "No News",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get news (V1)",
			"slug":    slug,
		})
	}
}
