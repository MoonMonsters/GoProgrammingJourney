package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// 令牌桶信息
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 令牌Bucket
type LimiterBucketRule struct {
	// 自定义键值对名称
	Key string
	// 间隔多久放N个令牌
	FillInterval time.Duration
	// 令牌桶的容量
	Capacity int64
	// 每次到达间隔时间后所放的具体令牌数量
	Quantum int64
}

// 接口
type LimiterIface interface {
	// 获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	// 获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// 新增令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}
