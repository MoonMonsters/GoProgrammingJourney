package main

import (
	"GoProgrammingJourney/blog_service/global"
	"GoProgrammingJourney/blog_service/internal/model"
	"GoProgrammingJourney/blog_service/internal/routers"
	"GoProgrammingJourney/blog_service/pkg/logger"
	"GoProgrammingJourney/blog_service/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go+Gin框架的博客项目
func main() {
	global.Logger.Info("blog_service.start")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = settings.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("JWT", &global.JWTSetting)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
