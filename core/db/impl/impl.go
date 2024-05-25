package impl

import (
	"feature-distributor/env"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type db interface {
	BuildDsn() string
	Build(dsn string) (dialector gorm.Dialector)
}

var (
	dbType = viper.GetString(env.DbType)
	dbUri  = viper.GetString(env.DbUri)
	dbUser = viper.GetString(env.DbUser)
	dbPass = viper.GetString(env.DbPass)
	dbHost = viper.GetString(env.DbHost)
	dbPort = viper.GetString(env.DbPort)
	dbName = viper.GetString(env.DbName)
)

func InitDb() (dialector gorm.Dialector) {
	var dsn string
	if dbUri != "" {
		dsn = dbUri
	}
	var d db
	switch dbType {
	case "postgres":
		panic("not implemented")
	default:
		d = &mysql{}
		if dbUri == "" {
			dsn = d.BuildDsn()
		}
		dialector = d.Build(dsn)
		break
	}
	return
}
