package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/db"
	"github.com/qingxunying/douyin/oss"
	"github.com/qingxunying/douyin/rdb"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	initConfig()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initConfig() {
	db.InitDb()
	rdb.InitRdb()
	oss.InitOss()

	// logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}
