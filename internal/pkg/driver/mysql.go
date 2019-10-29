package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

type DBMysqlOption struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	ConnectionSetting    ConnectionSetting
}

type ConnectionSetting struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewMysqlDatabase(option DBMysqlOption) *gorp.DbMap {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
	if err != nil {
		panic(fmt.Errorf("ERROR connect to DB MySQL: %s | %v", option.DBName, err))
	}

	db.SetConnMaxLifetime(option.ConnectionSetting.ConnMaxLifetime)
	db.SetMaxIdleConns(option.ConnectionSetting.MaxIdleConns)
	db.SetMaxOpenConns(option.ConnectionSetting.MaxOpenConns)

	gorp := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return gorp
}