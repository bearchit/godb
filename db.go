package godb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(opt ...interface{}) (*DB, error) {
	db, err := gorm.Open("mysql", opt...)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}
