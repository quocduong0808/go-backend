package middlewares

import (
	"fmt"
	"go/go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("---> authentication")
		token := ctx.GetHeader("Authorization")
		if token != "123456789" {
			response.ResponseError(ctx, response.ErrCodeTokenInvalid, response.MSG[response.ErrCodeTokenInvalid])
			ctx.Abort()
		}
		fmt.Println("---> authenticated")
		ctx.Next()
	}
}
