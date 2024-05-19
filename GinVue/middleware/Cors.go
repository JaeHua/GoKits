package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求,支持options访问

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 可将将 * 替换为指定的域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//放行的header也可以在后面自行添加
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//放行的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//永久有效
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			// 处理请求
			ctx.Next()
		}
	}
}
