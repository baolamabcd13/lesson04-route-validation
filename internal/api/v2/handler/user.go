package v2handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) GetUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of users (V2)",
	})
}

func (uh *UserHandler) GetUserByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V2)",
	})
}

func (uh *UserHandler) PostUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created (V2)",
	})
}

func (uh *UserHandler) PutUsersByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated (V2)",
	})
}

func (uh *UserHandler) DeleteUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "User deleted (V2)",
	})
}
