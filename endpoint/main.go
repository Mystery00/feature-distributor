package main

import (
	"feature-distributor/common/env"
	"feature-distributor/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	log := logrus.WithField("source", "main")

	address, exist := os.LookupEnv(env.EnvListenAddress)
	if !exist {
		address = ":7002"
	}
	logger.InitLog()
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		log.Infof(`EndpointServer is running at %s`, address)
		_ = srv.ListenAndServe()
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("EndpointServer exiting")
}
