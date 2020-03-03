package godb

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/creasty/defaults"

	"github.com/jinzhu/gorm"
)

type Configurer interface {
	Driver() string
	FormatDSN() string
}

type PostgresConfig struct {
	Host     string `validate:"required"`
	Port     int    `default:"5432"`
	DbName   string `validate:"required"`
	User     string `validate:"required"`
	Password string
	SSLMode  string `default:"disable"`
	LogMode  bool   `default:"false"`
}

func (c PostgresConfig) Driver() string {
	return "postgres"
}

func (c PostgresConfig) FormatDSN() string {
	args := make(map[string]interface{})
	args["host"] = c.Host
	args["port"] = c.Port
	args["user"] = c.User
	if c.Password != "" {
		args["password"] = c.Password
	}
	args["dbname"] = c.DbName
	args["sslmode"] = c.SSLMode

	s := make([]string, 0)
	for k, v := range args {
		s = append(s, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(s, " ")
}

func OpenGORM(configurer Configurer) (*gorm.DB, error) {
	if err := defaults.Set(configurer); err != nil {
		return nil, err
	}

	db, err := gorm.Open(configurer.Driver(), configurer.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return db, nil
}

type MySQLConfig struct {
	Addr      string `validate:"required"`
	DbName    string `validate:"required"`
	User      string `validate:"required"`
	Password  string `validate:"required"`
	ParseTime bool   `default:"true"`
}

func (c MySQLConfig) Driver() string {
	return "mysql"
}

func (c MySQLConfig) FormatDSN() string {
	mc := mysql.NewConfig()
	mc.Addr = c.Addr
	mc.DBName = c.DbName
	mc.User = c.User
	mc.Passwd = c.Password
	mc.ParseTime = c.ParseTime
	return mc.FormatDSN()
}
