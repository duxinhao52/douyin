package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qingxunying/douyin/conf"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func InitDb() {
	logrus.Infof("start init mysql...")
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("root:1234567890.@tcp(%s:5000)/douyin?charset=utf8&parseTime=True&loc=Local", conf.HostIp))
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
