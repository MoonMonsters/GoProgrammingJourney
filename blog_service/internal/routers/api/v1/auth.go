package v1

import (
	"GoProgrammingJourney/blog_service/global"
	"GoProgrammingJourney/blog_service/internal/service"
	"GoProgrammingJourney/blog_service/pkg/app"
	"GoProgrammingJourney/blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 通过appKey和appSecret, 获取token
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	global.Logger.Infof("GetAuth: valid: %v, errs: %v, param: %v", valid, errs, param)
	if valid != true {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))

		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)

		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)

		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
