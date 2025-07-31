package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Before -> Trước khi bắt đầu vào Handler
		log.Println("Start func - Check from middleware")
		ctx.Writer.Write([]byte("Start func - Check from middleware"))
		ctx.Next() //Đi vào -> Handler

		// After -> Sau khi Handler xử lí xong
		log.Println("End func - Check from middleware")
		ctx.Writer.Write([]byte("End func - Check from middleware"))

	}
}
