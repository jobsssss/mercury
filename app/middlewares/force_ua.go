// Package middlewares Gin 中间件
package middlewares

import (
	"errors"
	"mercury/pkg/response"

	"github.com/gin-gonic/gin"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取 User-Agent 标头信息
		if len(ctx.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(ctx, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}

		ctx.Next()
	}
}
