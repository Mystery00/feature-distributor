package db

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
)

type OrmLogger struct {
	*logrus.Entry
	logger.Config
}

func NewOrmLogger(log *logrus.Entry) logger.Interface {
	return &OrmLogger{
		Entry: log,
		Config: logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Warn,
		},
	}
}

// LogMode log mode
func (l *OrmLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l *OrmLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	log := l
	if l.LogLevel >= logger.Info {
		//去掉换行符
		m := strings.Replace(msg, "\n", "", -1)
		log.Infof(m, data...)
	}
}

// Warn print warn messages
func (l *OrmLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	log := l
	if l.LogLevel >= logger.Warn {
		//去掉换行符
		m := strings.Replace(msg, "\n", "", -1)
		log.Warnf(m, data...)
	}
}

// Error print error messages
func (l *OrmLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	log := l
	if l.LogLevel >= logger.Error {
		//去掉换行符
		m := strings.Replace(msg, "\n", "", -1)
		log.Errorf(m, data...)
	}
}

// Trace print sql message
func (l *OrmLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	log := l
	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		log.Errorf(`%s %s`, utils.FileWithLineNum(), err)
		if rows == -1 {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": "-",
			}).Error(sql)
		} else {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": rows,
			}).Error(sql)
		}
	default:
		if rows == -1 {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": "-",
			}).Debug(sql)
		} else {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": rows,
			}).Debug(sql)
		}
	}
}
