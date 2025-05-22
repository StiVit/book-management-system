package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("mysql", "stivit:vitalic12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}