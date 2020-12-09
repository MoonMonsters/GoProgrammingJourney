package global

import (
	"GoProgrammingJourney/blog_service/pkg/logger"
	"GoProgrammingJourney/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting      *setting.JWTSetting
	EmailSetting    *setting.EmailSetting
	Logger          *logger.Logger
)
