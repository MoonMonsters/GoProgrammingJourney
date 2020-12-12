package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// 上下文超时时间控制中间件
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 设置当前context的超时时间, 并重新赋值给gin.Context
		// 当当前请求运行到指定的时间后, 使用了该context的运行流程就会对context提供的超时时间做处理
		// 在指定的时间内取消请求
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
