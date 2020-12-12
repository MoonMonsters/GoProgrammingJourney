package routers

import (
	_ "GoProgrammingJourney/blog_service/docs"
	"GoProgrammingJourney/blog_service/global"
	"GoProgrammingJourney/blog_service/internal/middleware"
	v1 "GoProgrammingJourney/blog_service/internal/routers/api/v1"
	"GoProgrammingJourney/blog_service/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

// 限制了/auth的访问频率
// 限制时间间隔, 1秒钟
// 一秒钟内, 最多被访问10次
// 当一秒后, 重新放入10个令牌到令牌桶内, 也就是下一秒可再次被访问10次
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	// 令牌桶限制的url
	Key: "/auth",
	// 时间间隔
	FillInterval: time.Second * 1,
	// 令牌总容量
	Capacity: 10,
	// 重新放入令牌桶数量
	Quantum: 10,
})

func NewRouter() *gin.Engine {

	r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.AppInfo())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := v1.NewUpload()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	// 配置下载路径
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.GET("/auth", v1.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	apiV1.Use()
	{
		apiV1.POST("tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)

	}

	return r
}
