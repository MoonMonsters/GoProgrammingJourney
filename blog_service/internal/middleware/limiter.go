package middleware

import (
	"GoProgrammingJourney/blog_service/pkg/app"
	"GoProgrammingJourney/blog_service/pkg/errcode"
	"GoProgrammingJourney/blog_service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取令牌的key值
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// 传入1, 表示已使用一个令牌
			count := bucket.TakeAvailable(1)
			// 如果剩余可用令牌数为0, 则抛出异常, 禁止访问
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
