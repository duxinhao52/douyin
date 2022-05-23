package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/qingxunying/douyin/constdef"
	"github.com/sirupsen/logrus"
)

var ossClient *oss.Client

func InitOss() {
	logrus.Infof("[InitOss] start init oss...")
	client, err := oss.New(constdef.OssEndPoint, constdef.OssAccessKeyId, constdef.OssAccessKeySecret)
	if err != nil {
		logrus.Panicf("[InitOss] connect oss error, err=%+v", err)
	}
	ossClient = client
}
