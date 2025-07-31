package v1handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lesson04-route-validation/utils"
	"net/http"
	"os"
	"path/filepath"
)

type NewsHandler struct {
}

type PostNewsV1Param struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (nh *NewsHandler) GetNewsV1(ctx *gin.Context) {
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

func (nh *NewsHandler) PostNewsV1(ctx *gin.Context) {
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
		})
		return
	}

	//2<<20 -> 2^20 = 1048576 =  1MB
	if image.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "File is too large (5 MB)",
		})
		return
	}

	//os.ModePerm = 0777 -> read, write, excute for owner,group, other
	err = os.MkdirAll("./upload", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot create upload folder",
		})
	}
	//destination
	dst := fmt.Sprintf("./upload/%s", filepath.Base(image.Filename))

	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot save file",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post News (V1)",
		"name":    params.Title,
		"status":  params.Status,
		"image":   image.Filename,
		"path":    dst,
	})
}

func (nh *NewsHandler) PostUploadFileNewsV1(ctx *gin.Context) {
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
		})
		return
	}

	filename, err := utils.ValidateAndSaveFile(image, "./upload")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post News (V1)",
		"name":    params.Title,
		"status":  params.Status,
		"image":   filename,
		"path":    "./upload/" + filename,
	})
}

func (nh *NewsHandler) PostUploadMultipleFileNewsV1(ctx *gin.Context) {
	const publicURL = "http://localhost:8080/images/"
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid multipart form",
		})
		return
	}

	images := form.File["images"]
	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "No file provided",
		})
		return
	}

	var successFiles []string
	var failedFiles []map[string]string

	for _, image := range images {
		filename, err := utils.ValidateAndSaveFile(image, "./upload")
		if err != nil {
			failedFiles = append(failedFiles, map[string]string{
				"filename": image.Filename,
				"error":    err.Error(),
			})
			continue
		}
		publicImageURL := publicURL + filename
		successFiles = append(successFiles, publicImageURL)
	}

	resp := gin.H{
		"message":       "Post News (V1)",
		"name":          params.Title,
		"status":        params.Status,
		"success_files": successFiles,
	}

	if len(failedFiles) > 0 {
		resp["message"] = "Upload completed with partial errors"
		resp["errors_files"] = failedFiles
	}

	ctx.JSON(http.StatusOK, resp)
}
