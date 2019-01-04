package godb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}
