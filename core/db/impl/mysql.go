package impl

import (
	"fmt"
	driver "github.com/go-sql-driver/mysql"
	gormDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type mysql struct {
}

func (m *mysql) BuildDsn() string {
	hostAndPort := fmt.Sprintf("%s:%s", dbHost, dbPort)
	if dbPort == "" {
		hostAndPort = dbHost
	}
	conf := driver.NewConfig()
	conf.User = dbUser
	conf.Passwd = dbPass
	conf.Net = "tcp"
	conf.Addr = hostAndPort
	conf.DBName = dbName
	conf.ParseTime = true
	conf.Loc = time.Local
	conf.Params = make(map[string]string)
	conf.Params["charset"] = "utf8mb4"
	return conf.FormatDSN()
}

func (m *mysql) Build(dsn string) (dialector gorm.Dialector) {
	return gormDriver.Open(dsn)
}
