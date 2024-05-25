package db

import (
	"feature-distributor/core/db/impl"
	"feature-distributor/core/db/query"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitDatabase() {
	dialector := impl.InitDb()
	client, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
	}
	db, err := client.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	logrus.WithField("source", "main").Info("database connected")
	query.SetDefault(client)
}
