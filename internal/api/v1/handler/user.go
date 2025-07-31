package v1handler

import (
	"github.com/gin-gonic/gin"
	"lesson04-route-validation/utils"
	"log"

	"net/http"
)

type UserHandler struct{}

type GetUserByIdV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

type GetUserByUUIdV1Param struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of users (V1)",
	})
}

func (uh *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	var params GetUserByIdV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))

		return
	}
	log.Println("Into GetUserByIdV1")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User details (V1)",
		"user_id": params.ID,
	})
}

func (uh *UserHandler) GetUserByUUIdV1(ctx *gin.Context) {
	var params GetUserByUUIdV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "User details by UUID (V1)",
		"user_uuid": params.Uuid,
	})
}

func (uh *UserHandler) PostUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created (V1)",
	})
}

func (uh *UserHandler) PutUsersByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated (V1)",
	})
}

func (uh *UserHandler) DeleteUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "User deleted (V1)",
	})
}
