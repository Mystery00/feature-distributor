package env

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

func init() {
	configHome, exist := os.LookupEnv(EnvConfigHome)
	if !exist {
		//不存在环境变量，尝试取可执行文件目录下的 etc 目录
		configHome = "etc"
	}
	//判断配置文件目录是否存在
	if !exists(configHome) {
		//配置文件目录不存在，拒绝启动
		slog.Error("config file not exist")
		os.Exit(1)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configHome)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		slog.Info("config file updated")
		err := viper.ReadInConfig()
		if err != nil {
			slog.Error("read config failed", err)
		}
	})
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("read config failed", err)
		os.Exit(1)
	}
}

// exists 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	//获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
