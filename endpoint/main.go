package main

import (
	"context"
	"feature-distributor/common/env"
	"feature-distributor/common/logger"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/middleware"
	"feature-distributor/endpoint/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log := logrus.WithField("source", "main")

	address, exist := os.LookupEnv(env.EnvListenAddress)
	if !exist {
		address = ":7002"
	}
	logger.InitLog()
	grpc.Init()

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.ForwardedByClientIP = true
	middleware.SetMiddleware(router)
	web.Handle(router)

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

	if err := grpc.Close(); err != nil {
		log.Warnf("Close grpc connection: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("EndpointServer shutdown: %v", err)
	}
}
