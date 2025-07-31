package main

import (
	"github.com/gin-gonic/gin"
	v1handler "lesson03-route-group/internal/api/v1/handler"
	v2handler "lesson03-route-group/internal/api/v2/handler"
	"lesson03-route-group/utils"
)

func main() {
	r := gin.Default()

	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}

	userHandlerV1 := v1handler.NewUserHandler()

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUserByIdV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUserByUUIdV1)
			user.POST("/", userHandlerV1.PostUsersV1)
			user.PUT("/:id", userHandlerV1.PutUsersByIdV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersV1)
		}

		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()

			product.GET("/", productHandlerV1.GetProductsV1)
			product.GET("/:slug", productHandlerV1.GetProductBySLugV1)
			product.POST("/", productHandlerV1.PostProductsV1)
			product.PUT("/:id", productHandlerV1.PutProductsByIdV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductsV1)
		}

		category := v1.Group("/categories")
		{
			categoryHandlerV1 := v1handler.NewCategoryHandler()
			category.GET("/:category", categoryHandlerV1.GetCategoryBycategoryV1)
			category.POST("", categoryHandlerV1.PostCategoriesV1)
		}

		news := v1.Group("/news")
		{
			newsHandlerV1 := v1handler.NewNewsHandler()
			news.GET("/", newsHandlerV1.GetNewsV1)
			news.POST("/", newsHandlerV1.PostNewsV1)
			news.POST("/upload-file", newsHandlerV1.PostUploadFileNewsV1)
			news.GET("/:slug", newsHandlerV1.GetNewsV1)
		}

	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			userHandlerV2 := v2handler.NewUserHandler()

			user.GET("/", userHandlerV2.GetUsersV2)
			user.GET("/:id", userHandlerV2.GetUserByIdV2)
			user.POST("/", userHandlerV2.PostUsersV2)
			user.PUT("/:id", userHandlerV2.PutUsersByIdV2)
			user.DELETE("/:id", userHandlerV2.DeleteUsersV2)
		}
		product := v2.Group("/products")
		{
			productHandlerV2 := v2handler.NewProductHandler()

			product.GET("/", productHandlerV2.GetProductsV2)
			product.GET("/:id", productHandlerV2.GetproductByIdV2)
			product.POST("/", productHandlerV2.PostProductsV2)
			product.PUT("/:id", productHandlerV2.PutProductsByIdV2)
			product.DELETE("/:id", productHandlerV2.DeleteProductsV2)
		}
	}

	r.Run()
}
