package main

import (
	"feature-distributor/common/env"
	"feature-distributor/common/logger"
	"feature-distributor/core/db"
	"feature-distributor/core/grpc"
	"feature-distributor/core/provider"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	log := logrus.WithField("source", "main")

	address, exist := os.LookupEnv(env.EnvListenAddress)
	if !exist {
		address = ":7001"
	}
	logger.InitLog()
	db.InitDatabase()
	err := provider.Init()
	if err != nil {
		log.Fatalf("failed to init provider: %v", err)
	}

	go func() {
		log.Infof(`CoreServer is running at %s`, address)
		// 服务连接
		if err := grpc.Run(address); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("CoreServer exiting")
}
