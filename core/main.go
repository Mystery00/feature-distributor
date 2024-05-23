package main

import (
	"context"
	"errors"
	"feature-distributor/env"
	"feature-distributor/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log := logrus.WithField("source", "main")

	runPort, exist := os.LookupEnv(env.EnvRunPort)
	if !exist {
		runPort = "9090"
	}
	logger.InitLog()

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.ForwardedByClientIP = true
	//middleware.SetMiddleware(router)
	//health.Handle(router)
	//proxy.Handle(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(`:%s`, runPort),
		Handler: router,
	}

	go func() {
		log.Infof(`Server is running at :%s`, runPort)
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
