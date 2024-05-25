package logger

import (
	"feature-distributor/common/env"
	"fmt"
	"github.com/Mystery00/lumberjack"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func getLogFilePath(fileName string) string {
	logFileHome, exist := os.LookupEnv(env.EnvLogHome)
	if !exist {
		logFileHome = viper.GetString(env.LogHome)
	}
	err := os.MkdirAll(logFileHome, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}
	return fmt.Sprintf(`%s/%s`, logFileHome, fileName)
}

func InitLog() {
	local := viper.GetBool(env.LogLocal)
	showColor := false
	var out io.Writer
	if local {
		out = os.Stdout
		showColor = true
	} else {
		logFile := viper.GetString(env.LogFile)
		fileName := getLogFilePath(logFile)
		out = &lumberjack.Logger{
			Filename:         fileName,
			MaxSize:          256,
			MaxAge:           3,
			LocalTime:        true,
			Compress:         true,
			SplitByDay:       true,
			BackupTimeFormat: `2006-01-02.150405`,
		}
		showColor = viper.GetBool(env.LogColor)
	}
	//设置输出
	logrus.SetOutput(out)
	//设置日志级别
	logrus.SetLevel(logrus.InfoLevel)
	level := viper.GetString(env.LogLevel)
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        true,
		NoColors:        !showColor,
		ShowFullLevel:   false,
	})
	//添加钩子
	consoleLogger := logrus.New()
	consoleLogger.SetFormatter(logrus.StandardLogger().Formatter)
	if !local {
		logrus.AddHook(&logHook{
			logger: consoleLogger,
		})
	}
}

type logHook struct {
	logger *logrus.Logger
}

func (hook *logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *logHook) Fire(entry *logrus.Entry) error {
	source := entry.Data["source"]
	if source == "main" || entry.Level == logrus.PanicLevel || entry.Level == logrus.FatalLevel {
		//main的日志，往控制台打印一份
		hook.logger.WithFields(entry.Data).Logln(entry.Level, entry.Message)
	}
	return nil
}
