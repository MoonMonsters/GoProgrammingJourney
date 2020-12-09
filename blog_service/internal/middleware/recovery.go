package middleware

import (
	"GoProgrammingJourney/blog_service/global"
	"GoProgrammingJourney/blog_service/pkg/app"
	"GoProgrammingJourney/blog_service/pkg/email"
	"GoProgrammingJourney/blog_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {

		if err := recover(); err != nil {
			fmt.Printf("异常捕获: %v", err)
			global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
			err = defaultMailer.SendMail(
				global.EmailSetting.To,
				fmt.Sprintf("异常抛出, 发生时间: %d", time.Now().Unix()),
				fmt.Sprintf("错误信息: %v", err),
			)
			if err != nil {
				global.Logger.Panicf("mail.SendMail err: %v", err)
			}
			app.NewResponse(c).ToErrorResponse(errcode.ServerError)
			c.Abort()
		}
		c.Next()
	}
}
