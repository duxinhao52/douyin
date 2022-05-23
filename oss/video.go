package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/qingxunying/douyin/conf"
	"github.com/qingxunying/douyin/util"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"path/filepath"
)

func UpVideoToOss(file *multipart.FileHeader, userId int64) (string, error) {
	fileName := filepath.Base(file.Filename)
	finalName := fmt.Sprintf("%d_%s_%s", userId, util.CreateRandomString(1)[0], fileName)
	bucket, err := ossClient.Bucket(conf.VideoBucket)
	if err != nil {
		logrus.Errorf("[UpVideoToOss] get video bucket error, err=%+v", err)
		return "", err
	}
	data, _ := file.Open()
	err = bucket.PutObject(finalName, data, oss.ObjectACL(oss.ACLPublicRead))
	if err != nil {
		logrus.Errorf("[UpVideoToOss] upVideoToOss error, err=%+v", err)
		return "", err
	}

	videoUrl := conf.OssVideoUrlPrefix + finalName
	return videoUrl, nil
}
