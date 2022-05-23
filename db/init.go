package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func InitDb() {
	logrus.Infof("start init mysql...")
	var err error
	// 47.103.0.144:5000
	// 192.168.8.247:5000
	// localhost:3306
	db, err = gorm.Open("mysql", "root:1234567890.@tcp(47.103.0.144:5000)/douyin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Panicf("connect mysql error, err=%+v", err)
	}

	dbPool := db.DB()
	dbPool.SetMaxOpenConns(100)
	dbPool.SetMaxIdleConns(10)

	return
}

func GetDB() *gorm.DB {
	return db
}
